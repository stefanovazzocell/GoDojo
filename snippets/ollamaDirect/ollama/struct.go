package ollama

import (
	"encoding/base64"
)

// The response to a chat request to Ollama
//
// Source: https://github.com/ollama/ollama/blob/main/docs/api.md#response-8
type ChatResponse struct {
	// the model name
	Model string `json:"model,omitempty"`
	// the response message (if streaming, this would be a partial response)
	Message ChatMessage `json:"message"`
	// if true, the response is complete
	Done bool `json:"done"`
	// time spent generating the response
	TotalDurationNs uint64 `json:"total_duration"`
	// time spent in nanoseconds loading the model
	LoadDurationNs uint64 `json:"load_duration"`
	// number of tokens in the prompt
	PromptEvalTokens uint64 `json:"prompt_eval_count"`
	// time spent in nanoseconds evaluating the prompt
	PromptEvalDurationNs uint64 `json:"prompt_eval_duration"`
	// number of tokens the response
	ResponseTokens uint64 `json:"eval_count"`
	// time in nanoseconds spent generating the response
	ResponseDurationNs uint64 `json:"eval_duration"`
	// if present an error occurred
	Error string `json:"error,omitempty"`
}

// A ChatCompletion request
//
// Source: https://github.com/ollama/ollama/blob/main/docs/api.md#generate-a-chat-completion
type ChatRequest struct {
	// (required) the model name
	Model string `json:"model"`
	// the messages of the chat, this can be used to keep a chat memory
	Messages []ChatMessage `json:"messages"`
	// the format to return a response in. Currently the only accepted value is `json`
	Format string `json:"format,omitempty"`
	// additional model parameters listed in the documentation for the `Modelfile` such as `temperature`
	Options *ModelOptions `json:"options,omitempty"`
	// the prompt template to use (overrides what is defined in the `Modelfile`)
	Template string `json:"template,omitempty"`
	// if `false` the response will be returned as a single response object, rather than a stream of objects
	Stream bool `json:"stream,omitempty"`
	// Controls how long the model will stay loaded into memory following the request (default: 5m)
	KeepAlive string `json:"keep_alive,omitempty"`
}

// A Chat message
//
// Source: https://github.com/ollama/ollama/blob/main/docs/api.md#generate-a-chat-completion
type ChatMessage struct {
	// One of: system, user, or assistant
	Role string `json:"role"`
	// The message content
	Content string `json:"content"`
	// An optional list of images for multimodal models
	Images []string `json:"images,omitempty"`
}

// Appends an image to this chat message.
func (message *ChatMessage) AddImage(imageBytes []byte) {
	message.Images = append(message.Images, base64.StdEncoding.EncodeToString(imageBytes))
}

// Model options (i.e.: `Modelfile`)
//
// Source: https://github.com/ollama/ollama/blob/main/docs/modelfile.md#valid-parameters-and-values
type ModelOptions struct {
	// Enable Mirostat sampling for controlling perplexity. (default: 0, 0 = disabled, 1 = Mirostat, 2 = Mirostat 2.0)
	Mirostat int `json:"mirostat,omitempty"`
	// Influences how quickly the algorithm responds to feedback from the generated text. A lower learning rate will result in slower adjustments, while a higher learning rate will make the algorithm more responsive. (Default: 0.1)
	Mirostat_eta float64 `json:"mirostat_eta,omitempty"`
	// Controls the balance between coherence and diversity of the output. A lower value will result in more focused and coherent text. (Default: 5.0)
	Mirostat_tau float64 `json:"mirostat_tau,omitempty"`
	// Sets the size of the context window used to generate the next token. (Default: 2048)
	Num_ctx int `json:"num_ctx,omitempty"`
	// The number of GQA groups in the transformer layer. Required for some models, for example it is 8 for llama2:70b
	Num_gqa int `json:"num_gqa,omitempty"`
	// The number of layers to send to the GPU(s). On macOS it defaults to 1 to enable metal support, 0 to disable.
	Num_gpu int `json:"num_gpu,omitempty"`
	// Sets the number of threads to use during computation. By default, Ollama will detect this for optimal performance. It is recommended to set this value to the number of physical CPU cores your system has (as opposed to the logical number of cores).
	Num_thread int `json:"num_thread,omitempty"`
	// Sets how far back for the model to look back to prevent repetition. (Default: 64, 0 = disabled, -1 = num_ctx)
	Repeat_last_n int `json:"repeat_last_n,omitempty"`
	// Sets how strongly to penalize repetitions. A higher value (e.g., 1.5) will penalize repetitions more strongly, while a lower value (e.g., 0.9) will be more lenient. (Default: 1.1)
	Repeat_penalty float64 `json:"repeat_penalty,omitempty"`
	// The temperature of the model. Increasing the temperature will make the model answer more creatively. (Default: 0.8)
	Temperature float64 `json:"temperature,omitempty"`
	// Sets the random number seed to use for generation. Setting this to a specific number will make the model generate the same text for the same prompt. (Default: 0)
	Seed int `json:"seed,omitempty"`
	// Sets the stop sequences to use. When this pattern is encountered the LLM will stop generating text and return. Multiple stop patterns may be set by specifying multiple separate `stop` parameters in a modelfile.
	Stop string `json:"stop,omitempty"`
	// Tail free sampling is used to reduce the impact of less probable tokens from the output. A higher value (e.g., 2.0) will reduce the impact more, while a value of 1.0 disables this setting. (default: 1)
	Tfs_z float64 `json:"tfs_z,omitempty"`
	// Maximum number of tokens to predict when generating text. (Default: 128, -1 = infinite generation, -2 = fill context)
	Num_predict int `json:"num_predict,omitempty"`
	// Reduces the probability of generating nonsense. A higher value (e.g. 100) will give more diverse answers, while a lower value (e.g. 10) will be more conservative. (Default: 40)
	Top_k int `json:"top_k,omitempty"`
	// Works together with top-k. A higher value (e.g., 0.95) will lead to more diverse text, while a lower value (e.g., 0.5) will generate more focused and conservative text. (Default: 0.9)
	Top_p float64 `json:"top_p,omitempty"`
}
