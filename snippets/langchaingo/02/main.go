package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/schema"
)

func main() {
	// Setup the LLM
	llm, err := ollama.New(ollama.WithModel("gemma:2b"))
	if err != nil {
		panic(err)
	}

	// Setup

	var query string
	ctx := context.Background()
	now := time.Now()
	messages := []llms.MessageContent{
		llms.TextParts(schema.ChatMessageTypeSystem,
			fmt.Sprintf("You're a helpful assistant. To answer time-based queries you should know that today is %s %d %d.",
				now.Month().String(),
				now.Day(),
				now.Year())),
	}
	for {
		// Read the user query
		fmt.Print("User: ")
		reader := bufio.NewReader(os.Stdin)
		query, err = reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		// Append query to messages
		messages = append(messages, llms.TextParts(schema.ChatMessageTypeHuman, query))

		// Write reply
		fmt.Print("LLM: ")
		response, err := llm.GenerateContent(ctx, messages, llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
			// Print streamed reply
			fmt.Print(string(chunk))
			return nil
		}))
		if err != nil {
			panic(err)
		}
		fmt.Println()

		// Append answer to messages
		if len(response.Choices) > 0 {
			messages = append(messages, llms.TextParts(schema.ChatMessageTypeAI, response.Choices[0].Content))
		}
	}
}
