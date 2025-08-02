/*
Copyright © 2025 Abinand P <abinand0911@gmail.com>
*/
package gemini

import (
	"encoding/json"

	"github.com/Abiji-2020/PesudoCLI/internal/redisclient"
	"google.golang.org/genai"
)

func (g *GeminiClient) AskQuestion(question, model string, contextValues []redisclient.QuerySearchResult) (*GeminiResponse, error) {

	config := &genai.GenerateContentConfig{
		ResponseMIMEType: "application/json",
		ResponseJsonSchema: &genai.Schema{
			Type: "object",
			Properties: map[string]*genai.Schema{
				"command":     {Type: "string"},
				"os":          {Type: "string"},
				"explanation": {Type: "string"},
				"answer":      {Type: "string"},
			},
			Required: []string{"command", "os", "explanation", "answer"},
		},
		Temperature: func() *float32 { v := float32(0.15); return &v }(),
	}
	chat, err := g.client.Chats.Create(g.ctx, model, config, []*genai.Content{genai.NewContentFromText(BuildContextPrompt(contextValues), genai.RoleUser)})
	if err != nil {
		return nil, err
	}
	result, err := chat.SendMessage(g.ctx, genai.Part{Text: question + "\n Answer in the following  with  the context to explain clearly"})
	if err != nil {
		return nil, err
	}
	var parsedValues GeminiResponse
	err = json.Unmarshal([]byte(result.Text()), &parsedValues)
	if err != nil {
		return nil, err
	}
	return &GeminiResponse{
		Answer:      parsedValues.Answer,
		Command:     parsedValues.Command,
		Os:          parsedValues.Os,
		Explanation: parsedValues.Explanation,
	}, nil
}
