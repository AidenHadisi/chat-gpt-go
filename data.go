package chatgptgo

import "fmt"

// Request is the request body for the chat endpoint.
type Request struct {
	// Model is ID of the model to use. Currently, only `gpt-3.5-turbo` and `gpt-3.5-turbo-0301` are supported.
	Model Model `json:"model"`

	// Messages is the messages to generate chat completions for, in the chat format.
	Messages []*Message `json:"messages"`

	//Temperature is what sampling temperature to use, between 0 and 2. Higher values like 0.8 will make the output more random, while lower values like 0.2 will make it more focused and deterministic.
	Temperature float64 `json:"temperature,omitempty"`

	// TopP is an alternative to sampling with temperature, called nucleus sampling, where the model considers the results of the tokens with top_p probability mass.
	TopP float64 `json:"top_p,omitempty"`

	// N is how many chat completion choices to generate for each input message.
	N int `json:"n,omitempty"`

	//Stop is up to 4 sequences where the API will stop generating further tokens.
	Stop []string `json:"stop,omitempty"`

	// MaxTokens is the maximum number of tokens to generate for each chat completion choice.
	MaxTokens int `json:"max_tokens,omitempty"`

	// PresencePenalty is a number between -2.0 and 2.0. Positive values penalize new tokens based on whether they appear in the text so far, increasing the model's likelihood to talk about new topics.
	PresencePenalty float64 `json:"presence_penalty,omitempty"`

	// FrequencyPenalty is a number between -2.0 and 2.0. Positive values penalize new tokens based on their existing frequency in the text so far, decreasing the model's likelihood to repeat the same line verbatim.
	FrequencyPenalty float64 `json:"frequency_penalty,omitempty"`

	// User is a unique identifier representing your end-user, which can help OpenAI to monitor and detect abuse.
	User string `json:"user,omitempty"`
}

// Response is the response body returned from the chat endpoint.
type Response struct {
	ID       string    `json:"id"`
	Object   string    `json:"object"`
	Created  int64     `json:"created"`
	Choices  []*Choice `json:"choices"`
	Usage    *Usage    `json:"usage"`
	ThreadId string    `json:"-"`
}

// Choice is a single choice (message) returned from the chat endpoint.
type Choice struct {
	Index        int      `json:"index"`
	Message      *Message `json:"message"`
	FinishReason string   `json:"finish_reason"`
}

// Message is a single message sent to and returned from the chat endpoint.
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// Usage is the usage information returned from the chat endpoint.
type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type ApiError struct {
	ErrorDetails *ErrorDetails `json:"error"`
}

func (e *ApiError) Error() string {
	return fmt.Sprintf("ChatGPT API Error: %s", e.ErrorDetails.Message)
}

type ErrorDetails struct {
	Message string `json:"message"`
	Type    string `json:"type"`
	Code    string `json:"code"`
}
