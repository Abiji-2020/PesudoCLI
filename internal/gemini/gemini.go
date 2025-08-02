/*
Copyright © 2025 Abinand P <abinand0911@gmail.com>
*/
package gemini

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/Abiji-2020/PesudoCLI/internal/redisclient"
	"github.com/Abiji-2020/PesudoCLI/pkg/io"
	"google.golang.org/genai"
)

type GeminiClient struct {
	client *genai.Client
	ctx    context.Context
}

type Embedder interface {
	Embed(docs []io.CommandDoc, model string) ([]float32, error)
	EmbedQuestion(question string, model string) ([]float32, error)
}

type GeminiChatInterface interface {
	AskQuestion(question string, model string, contextValues []redisclient.QuerySearchResult) (string, error)
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

	var contents []*genai.Content
	for _, doc := range docs {
		content := genai.NewContentFromText(doc.TextChunk, genai.RoleUser)
		contents = append(contents, content)
	}

	var embeddedDocs []io.CommandDoc
	for i := 0; i < len(contents); i += 20 {
		end := i + 20
		if end > len(contents) {
			end = len(contents)
		}
		sample := contents[i:end]

		embedding, err := g.client.Models.EmbedContent(
			g.ctx, model, sample, &genai.EmbedContentConfig{
				TaskType:             "QUESTION_ANSWERING",
				OutputDimensionality: func(i int32) *int32 { return &i }(3072),
			})
		if err != nil {
			if strings.Contains(err.Error(), "Resource has been exhausted") {
				fmt.Println("Rate limit exceeded, adding the values to redis")

				return embeddedDocs, nil
			}
			return nil, fmt.Errorf("failed to embed text: %w", err)
		}
		if len(embedding.Embeddings) == 0 {
			return nil, fmt.Errorf("no embeddings returned for the batch of text")
		}
		fmt.Println("Embedding response received for batch:", i/20+1)
		embeddings := embedding.Embeddings
		temp := i
		for j := range embeddings {
			if len(embeddings[j].Values) == 0 {
				return nil, fmt.Errorf("no embedding values returned for text: %s", docs[i].TextChunk)
			}

			docs[temp].Embedding = embedding.Embeddings[j].Values
			embeddedDocs = append(embeddedDocs, docs[temp])
			temp++
		}
		time.Sleep(10 * time.Second) // Sleep to avoid rate limiting
	}
	return embeddedDocs, nil

}

func (g *GeminiClient) EmbedQuestion(question string, model string) ([]float32, error) {
	embedding, err := g.client.Models.EmbedContent(
		g.ctx, model, []*genai.Content{genai.NewContentFromText(question, genai.RoleUser)}, &genai.EmbedContentConfig{
			TaskType:             "QUESTION_ANSWERING",
			OutputDimensionality: func(i int32) *int32 { return &i }(3072), // Adjust as needed
		})
	if err != nil {
		return nil, fmt.Errorf("failed to embed question: %w", err)
	}
	if len(embedding.Embeddings) == 0 || len(embedding.Embeddings[0].Values) == 0 {
		return nil, fmt.Errorf("no embedding values returned for question: %s", question)
	}
	return embedding.Embeddings[0].Values, nil
}
