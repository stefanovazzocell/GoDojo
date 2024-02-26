package ollama

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const (
	// The ollama chat completion api path
	chatCompletionApiPath = "/api/chat"
)

// Creates a chat completion request and executes it as a stream.
//
// Requires the chat messages (history + last message) and a ResponseFn which will be called when we receive bytes
func (client *OllamaClient) ChatCompletionStream(Messages []ChatMessage, ResponseFn func(chunk []byte)) (chatResponse *ChatResponse, err error) {
	// Create request body
	requestBytes, err := json.Marshal(client.getChatRequest(Messages))
	if err != nil {
		err = fmt.Errorf("failed to marshal chat request: %w", err)
		return
	}

	// Setup request
	request, err := http.NewRequest(http.MethodPost, client.ServerUrl+chatCompletionApiPath, bytes.NewReader(requestBytes))
	if err != nil {
		err = fmt.Errorf("failed to create request body: %w", err)
		return
	}
	// Set request header
	request.Header = client.DefaultHeader

	// Execute request
	response, err := client.HttpClient.Do(request)
	if err != nil {
		err = fmt.Errorf("failed to do request: %w", err)
		return
	}
	defer response.Body.Close()

	// Prepare to stream response
	decoder := json.NewDecoder(response.Body)
	var chatMessage *ChatMessage = nil
	chatResponse = &ChatResponse{}
	for {
		// Decode the next part
		err = decoder.Decode(chatResponse)
		if err == io.EOF {
			err = errors.New("got EOF from ollama before being done")
			return
		}
		if err != nil {
			err = fmt.Errorf("response decoding failed: %w", err)
			return
		}

		// Check if we got an error back
		if len(chatResponse.Error) > 0 {
			err = errors.New("ollama returned the following error: " + chatResponse.Error)
			return
		}

		// Check if we're done
		if chatResponse.Done {
			// Append the full message
			chatResponse.Message = *chatMessage
			break
		}

		// Return response
		ResponseFn([]byte(chatResponse.Message.Content))

		// Append response to chatMessage
		if chatMessage == nil {
			// First chunk
			chatMessage = &ChatMessage{
				Role:    chatResponse.Message.Role,
				Content: chatResponse.Message.Content,
			}
		} else {
			chatMessage.Content = chatMessage.Content + chatResponse.Message.Content
		}
	}

	return
}

// Returns a chat request for this client and message logs
func (client *OllamaClient) getChatRequest(Messages []ChatMessage) *ChatRequest {
	return &ChatRequest{
		Model:    client.Model,
		Messages: Messages,
		Options:  client.DefaultOptions,
		Stream:   true,
	}
}
