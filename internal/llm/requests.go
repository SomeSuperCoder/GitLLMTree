package llm

import (
	"context"

	"github.com/sashabaranov/go-openai"
)

func GetStandardChatCompletion(ctx context.Context, client *openai.Client) (openai.ChatCompletionResponse, error) {
	return client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: openai.GPT4oMini,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Hello!",
				},
			},
		},
	)
}

func ChatCompletionToString(c openai.ChatCompletionResponse) string {
	return c.Choices[0].Message.Content
}
