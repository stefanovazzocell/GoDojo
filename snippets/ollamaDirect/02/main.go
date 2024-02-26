package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/stefanovazzocell/GoDojo/snippets/ollamaDirect/ollama"
)

var (
	// Source: https://github.com/langchain-ai/langchain/blob/master/libs/experimental/langchain_experimental/llms/ollama_functions.py
	DefaultSystemTemplate = `You are a helpful assistant and have access to the following tools:

%s

You must always select one of the above tools and respond with only a JSON object matching the following schema:

{{
  "tool": <name of the selected tool>,
  "tool_input": <parameters for the selected tool, matching the tool's JSON schema>
}}

If you don't know the answer, don't make it up. Use the tools to help you answering when appropriate.`
	// Source: https://github.com/langchain-ai/langchain/blob/master/libs/experimental/langchain_experimental/llms/ollama_functions.py
	DefaultResponseSchema = `{
	"name": "__conversational_response",
	"description": (
		"Respond conversationally if no other tools should be called for a given query."
	),
	"parameters": {
		"type": "object",
		"properties": {
			"response": {
				"type": "string",
				"description": "Conversational response to the user.",
			},
		},
		"required": ["response"],
	},
}`
	// Source: https://github.com/langchain-ai/langchain/blob/master/libs/experimental/tests/integration_tests/llms/test_ollama_functions.py
	MockWeatherResponseSchema = `{
		"name": "get_current_weather",
		"description": (
			"Get the current weather in a given location"
		),
		"parameters": {
			"type": "object",
			"properties": {
				"location": {
					"type": "string",
					"description": "The city and state, e.g. San Francisco, CA",
				},
				"unit": {
					"type": "string",
					"enum": ["celsius", "fahrenheit"],
				}
			},
			"required": ["location"],
		},
	}`
)

func main() {
	// Setup the LLM
	// - mistral
	// - gemma:2b
	// - llama2:7b
	llm := ollama.NewOllamaClient("http://localhost:11434", "mistral", nil)

	// Setup the tools
	tools := []string{DefaultResponseSchema, MockWeatherResponseSchema}

	// Setup the chat
	var query string = fmt.Sprintf(DefaultSystemTemplate, strings.Join(tools, "\n"))
	var err error
	messages := []ollama.ChatMessage{
		{
			Role:    "system",
			Content: query,
		},
	}
	fmt.Println("------\nSystem: " + query + "\n")

	// Chat loop
	var fnCall bool = false
	for {
		if fnCall {
			// We're simulating a function call
			query = `The function "get_current_weather" returned the following: "25 degrees and sunny".`
			fmt.Println("------\nSystem: " + query + "\n")

			// Append query to messages
			messages = append(messages, ollama.ChatMessage{Role: "system", Content: query})
		} else {
			// Read the user query
			fmt.Print("------\nUser: ")
			reader := bufio.NewReader(os.Stdin)
			query, err = reader.ReadString('\n')
			if err != nil {
				panic(err)
			}
			fmt.Println()

			// Append query to messages
			messages = append(messages, ollama.ChatMessage{Role: "user", Content: query})
		}

		// Write reply
		fmt.Print("------\nLLM: ")
		fullResponse, err := llm.ChatCompletionStream(messages, func(chunk []byte) {
			// Print streamed reply
			fmt.Print(string(chunk))
		})
		if err != nil {
			panic(err)
		}
		fmt.Print("\n\n")

		// Check if call to get_current_weather
		fnCall = strings.Contains(fullResponse.Message.Content, "get_current_weather")
		// Append answer to messages
		messages = append(messages, fullResponse.Message)
	}
}
