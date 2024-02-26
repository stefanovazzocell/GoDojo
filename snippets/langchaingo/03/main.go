package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/tmc/langchaingo/agents"
	"github.com/tmc/langchaingo/callbacks"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/memory"
	"github.com/tmc/langchaingo/tools"
)

func main() {
	// Setup the LLM
	// llama2:7b, gemma:2b, mistral
	llm, err := ollama.New(ollama.WithModel("mistral"))
	if err != nil {
		panic(err)
	}

	// Setup tools
	agentTools := []tools.Tool{
		NewWeatherForecastMock(),
		tools.Calculator{
			CallbacksHandler: callbacks.LogHandler{},
		},
	}

	// Setup memory
	mem := memory.NewConversationBuffer()

	// Setup agent
	agent, err := agents.Initialize(llm, agentTools, agents.ConversationalReactDescription, agents.WithMaxIterations(6), agents.WithMemory(mem), agents.WithCallbacksHandler(callbacks.LogHandler{}))
	if err != nil {
		panic(err)
	}

	var query string
	ctx := context.Background()
	for {
		// Read the user query
		fmt.Print("User: ")
		reader := bufio.NewReader(os.Stdin)
		query, err = reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		// Write reply
		fmt.Print("LLM: ")
		output, err := agent.Call(ctx, map[string]any{
			"input": query,
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(output["output"])
		fmt.Println()
	}
}
