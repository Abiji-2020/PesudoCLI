/*
Copyright © 2025 Abinand P <abinand0911@gmail.com>
*/
package gemini

import (
	"context"
	"fmt"

	"github.com/Abiji-2020/PesudoCLI/pkg/io"
	"google.golang.org/genai"
)

type GeminiClient struct {
	client *genai.Client
	ctx    context.Context
}

type Embedder interface {
	Embed(text string) ([]float32, error)
}

func NewGeminiClient(apiKey string) (*GeminiClient, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create Gemini client: %w", err)
	}
	return &GeminiClient{
		client: client,
		ctx:    ctx,
	}, nil
}

func (g *GeminiClient) Embed(docs []io.CommandDoc, model string) ([]io.CommandDoc, error) {
	var embeddedDocs []io.CommandDoc
	for _, doc := range docs {
		embedding, err := g.client.Models.EmbedContent(
			g.ctx, model, genai.Text(doc.TextChunk),
			&genai.EmbedContentConfig{
				TaskType: "RETRIEVAL_QUERY",
			})
		if err != nil {
			return nil, fmt.Errorf("failed to embed text: %w", err)
		}
		if len(embedding.Embeddings) == 0 {
			return nil, fmt.Errorf("no embeddings returned for text: %s", doc.TextChunk)
		}
		doc.Embedding = embedding.Embeddings[0].Values
		embeddedDocs = append(embeddedDocs, doc)
	}
	return embeddedDocs, nil

}
