package main

import (
	"context"
	"log"

	"github.com/SomeSuperCoder/GitLLMTree/config"
	"github.com/SomeSuperCoder/GitLLMTree/internal/llm"
)

func main() {
	ctx := context.Background()

	appConfig := config.LoadConfig()
	client, err := llm.NewClient(appConfig)
	if err != nil {
		log.Panic(err)
	}

	resp, err := llm.GetStandardChatCompletion(ctx, client)
	if err != nil {
		log.Panic(err)
	}

	responseString := llm.ChatCompletionToString(resp)
	log.Println(responseString)
}
