package main

import (
	"context"
	"fmt"
	"log"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"

	"io/ioutil"

	"github.com/lukegriffith/midori/pkg/journal"
)

func main() {
	llm, err := ollama.New(ollama.WithModel("llama2"))
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()

	content, err := ioutil.ReadFile("prompt.txt")
	if err != nil {
		log.Fatal("prompt.txt does not exist", err)
	}

	jContent, err := journal.ListJournal()

	if err != nil {
		log.Fatal(err)
	}

	prompt := fmt.Sprintf(string(content), jContent)

	fmt.Println(prompt)

	completion, err := llms.GenerateFromSinglePrompt(
		ctx,
		llm,
		prompt,
		llms.WithTemperature(0.8),
		llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
			fmt.Print(string(chunk))
			return nil
		}),
	)
	if err != nil {
		log.Fatal(err)
	}

	_ = completion
}
