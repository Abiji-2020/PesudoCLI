/*
Copyright © 2025 Abinand P <abinand0911@gmail.com>
*/

package redisclient

import (
	"context"
	"fmt"
	"log"

	"github.com/Abiji-2020/PesudoCLI/pkg/io"
	"github.com/Abiji-2020/PesudoCLI/pkg/utils"
	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	ctx    context.Context
	client *redis.Client
}

type RedisClientInterface interface {
	Client() *redis.Client
	Close() error
	Context() context.Context
}

func NewRedisClient(addr string) *RedisClient {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	if err := client.Ping(ctx).Err(); err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}

	return &RedisClient{
		ctx:    ctx,
		client: client,
	}
}

func (r *RedisClient) Client() *redis.Client {
	return r.client
}

func (r *RedisClient) Close() error {
	if err := r.client.Close(); err != nil {
		log.Printf("Error closing Redis client: %v", err)
		return err
	}
	return nil
}

func (r *RedisClient) Context() context.Context {
	return r.ctx
}

func (r *RedisClient) CreateVectorIndex(indexName string, dim int) error {
	_, err := r.client.Do(r.ctx, "FT.INFO", indexName).Result()
	if err == nil {
		log.Printf("Index %s already exists", indexName)
		return nil
	}

	commandToRun := []interface{}{
		"FT.CREATE", indexName,
		"ON", "HASH",
		"PREFIX", "1", "doc:",
		"SCHEMA",
		"command", "TEXT",
		"os", "TEXT",
		"text_chunk", "TEXT",
		"embedding", "VECTOR", "HNSW", "6",
		"DIM", dim,
		"TYPE", "FLOAT32",
		"DISTANCE_METRIC", "COSINE",
	}
	_, err = r.client.Do(r.ctx, commandToRun...).Result()

	if err != nil {
		log.Printf("Error creating index %s: %v", indexName, err)
		return err
	}

	return nil
}

func (r *RedisClient) AddDocument(docs []io.CommandDoc) error {
	var errors []error
	for _, doc := range docs {
		id := utils.GetID(doc.Command, doc.Os)
		bytesEmbedding, err := utils.Float32SliceToBytes(doc.Embedding)
		if err != nil {
			log.Fatal("failed to convert embedding to bytes: %w", err)
			errors = append(errors, err)
			continue
		}
		key := "doc:" + id

		fields := map[string]interface{}{
			"command":    doc.Command,
			"os":         doc.Os,
			"text_chunk": doc.TextChunk,
			"embedding":  bytesEmbedding,
		}
		if err := r.client.HSet(r.ctx, key, fields).Err(); err != nil {
			log.Fatal("failed to add document to Redis: %w", err)
			errors = append(errors, err)
			continue
		}
	}
	if len(errors) > 0 {
		return fmt.Errorf("encountered errors while adding documents: %v", errors)
	}
	log.Printf("Added %d documents to Redis", len(docs))
	return nil
}
