package main

import (
	"log"

	"github.com/SomeSuperCoder/GitLLMTree/config"
	"github.com/SomeSuperCoder/GitLLMTree/internal"
)

func main() {
	appConfig := config.LoadConfig()

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
	log.Println(diff)
}
