package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

func main() {
	// Setup the LLM
	llm, err := ollama.New(ollama.WithModel("gemma:2b"))
	if err != nil {
		panic(err)
	}

	// Read the user query
	fmt.Print("User: ")
	reader := bufio.NewReader(os.Stdin)
	query, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	// Write reply
	ctx := context.Background()
	fmt.Print("LLM: ")
	_, err = llms.GenerateFromSinglePrompt(ctx, llm, query,
		llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
			// Print streamed reply
			fmt.Print(string(chunk))
			return nil
		}))
	if err != nil {
		panic(err)
	}
	fmt.Println()
}
