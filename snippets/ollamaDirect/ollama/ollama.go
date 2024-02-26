package ollama

import (
	"fmt"
	"net/http"
	"runtime"
	"time"
)

const (
	// The default timeout for the client
	defaultClientTimeout = 5 * time.Minute
)

// A Ollama client
type OllamaClient struct {
	// The default http client
	HttpClient http.Client
	// Ollama server url
	ServerUrl string
	// The model to use
	Model string
	// Default model options to use
	DefaultOptions *ModelOptions
	// User Agent for requests
	DefaultHeader http.Header
}

// Returns a new OllamaClient with the specified parameters set
//
// If tempearture is nil, the default will be used
func NewOllamaClient(ServerUrl string, Model string, Temperature *float64) *OllamaClient {
	// Setup client
	defaultClient := http.Client{
		Timeout: defaultClientTimeout,
	}

	// Setup header
	defaultHeader := http.Header{}
	defaultHeader.Set("Content-Type", "application/json")
	defaultHeader.Set("Accept", "application/x-ndjson")
	defaultHeader.Set("User-Agent", fmt.Sprintf("ollama-cli (%s %s) Go/%s", runtime.GOARCH, runtime.GOOS, runtime.Version()))

	// Setup Options if any
	var defaultOptions *ModelOptions = nil
	if Temperature != nil {
		defaultOptions = &ModelOptions{
			Temperature: *Temperature,
		}
	}

	// Return struct
	return &OllamaClient{
		HttpClient:     defaultClient,
		ServerUrl:      ServerUrl,
		Model:          Model,
		DefaultOptions: defaultOptions,
		DefaultHeader:  defaultHeader,
	}
}
