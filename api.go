// Package chatgptgo provides a client for communicating with the OpenAI's GPT-3.5 (ChatGPT) API.
package chatgptgo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	// BaseUrl is the base URL for the OpenAI API.
	BaseUrl = "https://api.openai.com/v1"

	// ChatUrl is the URL for the chat endpoint.
	ChatUrl = BaseUrl + "/chat/completions"
)

// Api is the client for communicating with the OpenAI API.
type Api struct {
	organizationId string
	key            string
	client         *http.Client
}

// NewApi creates a new Api instance.
func NewApi(key string) *Api {
	return &Api{
		key:    key,
		client: http.DefaultClient,
	}
}

// WithClient sets your custom HTTP client to use for requests.
func (a *Api) WithClient(client *http.Client) *Api {
	a.client = client
	return a
}

// WithOrganizationId sets your organization ID to use for requests.
func (a *Api) WithOrganizationId(organizationId string) *Api {
	a.organizationId = organizationId
	return a
}

// Chat sends a chat request to the OpenAI API and returns the response.
func (a *Api) Chat(r *Request) (*Response, error) {
	body, err := json.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, ChatUrl, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	response, err := a.do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	if response.StatusCode != http.StatusOK {
		apiError := &ApiError{
			StatusCode: response.StatusCode,
		}
		if err := json.NewDecoder(response.Body).Decode(apiError); err != nil {
			return nil, fmt.Errorf("request failed with status code %d", response.StatusCode)
		}

		return nil, apiError
	}

	result := &Response{}
	if err := json.NewDecoder(response.Body).Decode(result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return result, nil
}

func (a *Api) do(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", "Bearer "+a.key)
	req.Header.Set("Content-Type", "application/json")
	if a.organizationId != "" {
		req.Header.Set("OpenAI-Organization", a.organizationId)
	}

	return a.client.Do(req)
}
