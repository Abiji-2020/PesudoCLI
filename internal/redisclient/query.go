/*
Copyright © 2025 Abinand P <abinand0911@gmail.com>
*/
package redisclient

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/Abiji-2020/PesudoCLI/pkg/utils"
	"github.com/redis/go-redis/v9"
)

func (r *RedisClient) QuerySearch(indexName string, topK string, query []float32) ([]QuerySearchResult, error) {
	vectorbytes, err := utils.Float32SliceToBytes(query)
	if err != nil {
		return nil, err
	}
	err = utils.WriteVectorToFile(query, "query_vector.bin")
	if err != nil {
		return nil, fmt.Errorf("error writing query vector to file: %w", err)
	}
	result, err := r.client.FTSearchWithArgs(r.ctx, indexName, "*=>[KNN "+topK+" @embedding $vectorbytes as vectoranswer]",
		&redis.FTSearchOptions{
			Return: []redis.FTSearchReturn{
				{FieldName: "vectoranswer"},
				{FieldName: "text_chunk"},
				{FieldName: "os"},
				{FieldName: "command"},
			},
			Params: map[string]any{
				"vectorbytes": vectorbytes,
			},
			DialectVersion: 2,
			SortBy: []redis.FTSearchSortBy{
				{FieldName: "vectoranswer"},
			},
		}).RawResult()

	if err != nil {
		return nil, fmt.Errorf("error querying Redis: %w", err)

	}

	cleaned := utils.ConvertInterfaceMap(result)

	jsonBytes, err := json.MarshalIndent(cleaned, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal: %v", err)
	}

	var searchResult RedisSearchResult
	if err := json.Unmarshal(jsonBytes, &searchResult); err != nil {
		log.Fatalf("Failed to unmarshal: %v", err)
	}
	var results []QuerySearchResult
	for _, doc := range searchResult.Results {
		vectorDistance, err := strconv.ParseFloat(doc.ExtraAttributes.VectorAnswer, 64)
		if err != nil {
			log.Printf("Error parsing vector distance: %v", err)
			continue
		}
		results = append(results, QuerySearchResult{
			Command:        doc.ExtraAttributes.Command,
			Os:             doc.ExtraAttributes.Os,
			TextChunk:      doc.ExtraAttributes.TextChunk,
			VectorDistance: vectorDistance,
		})

	}
	return results, nil
}
