package llm

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/sashabaranov/go-openai"
)

func makeChatCompletionRequest(prompt string) openai.ChatCompletionRequest {
	return openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
	}
}

func GetStandardChatCompletion(ctx context.Context, client *openai.Client, prompt string) (openai.ChatCompletionResponse, error) {
	return client.CreateChatCompletion(ctx, makeChatCompletionRequest(prompt))
}

func ChatCompletionResponseToString(c openai.ChatCompletionResponse) string {
	return c.Choices[0].Message.Content
}

func GetStreamChatCompletion(ctx context.Context, client *openai.Client, prompt string) (*openai.ChatCompletionStream, error) {
	return client.CreateChatCompletionStream(ctx, makeChatCompletionRequest(prompt))
}

func LogChatCompletionStream(stream *openai.ChatCompletionStream) error {
	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println()
			return nil
		}

		if err != nil {
			return err
		}

		fmt.Print(response.Choices[0].Delta.Content)
	}
}
