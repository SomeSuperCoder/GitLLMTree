package main

import (
	"context"
	"log"

	"github.com/SomeSuperCoder/GitLLMTree/config"
	"github.com/SomeSuperCoder/GitLLMTree/internal"
	"github.com/SomeSuperCoder/GitLLMTree/internal/llm"
	"github.com/SomeSuperCoder/GitLLMTree/internal/prompts"
)

func main() {
	ctx := context.Background()

	appConfig := config.LoadConfig()
	client, err := llm.NewClient(appConfig)
	if err != nil {
		log.Panic(err)
	}

	// Get the checkpoint diff
	commitPair := &internal.CommitPair{
		Owner:        "SomeSuperCoder",
		Repo:         "NeonRust",
		BaseCommit:   "0138d6ded61beb8ebc3e4f4dfcd27467569c12da",
		TargetCommit: "f53e9eb5ddf45ac43ef83ae6af37af03a720207b",
	}

	diff, err := internal.RequestDiff(commitPair, appConfig.GithubToken)
	if err != nil {
		log.Panicln(err)
	}

	// Form the LLM prompt
	prompt := prompts.DiffSummary([]string{}, diff, "Create a Solana-like blochchain")

	// Get the LLM summary
	resp, err := llm.GetStandardChatCompletion(ctx, client, prompt)
	if err != nil {
		log.Panic(err)
	}

	responseString := llm.ChatCompletionToString(resp)
	log.Println(responseString)
}
