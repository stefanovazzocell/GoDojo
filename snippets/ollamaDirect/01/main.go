package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/stefanovazzocell/GoDojo/snippets/ollamaDirect/ollama"
)

func main() {
	// Setup the LLM
	// - mistral
	// - gemma:2b
	// - llama2:7b
	llm := ollama.NewOllamaClient("http://localhost:11434", "mistral", nil)

	// Setup the chat
	var query string
	var err error
	messages := []ollama.ChatMessage{
		{
			Role:    "system",
			Content: "You're a helpful assistant.",
		},
	}

	// Chat loop
	for {
		// Read the user query
		fmt.Print("User: ")
		reader := bufio.NewReader(os.Stdin)
		query, err = reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		// Append query to messages
		messages = append(messages, ollama.ChatMessage{Role: "user", Content: query})

		// Write reply
		fmt.Print("LLM: ")
		fullResponse, err := llm.ChatCompletionStream(messages, func(chunk []byte) {
			// Print streamed reply
			fmt.Print(string(chunk))
		})
		if err != nil {
			panic(err)
		}
		fmt.Println()

		// Append answer to messages
		messages = append(messages, fullResponse.Message)
	}
}
