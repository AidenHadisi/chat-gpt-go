# ChatGPT Go

Communicate with OpenAi's GPT3.5 (ChatGPT) API.

## Usage
    
```go
    package main

    import chatgptgo "github.com/AidenHadisi/chat-gpt-go"

    func main() {
        api := chatgptgo.NewApi("YOUR_API_KEY")

        request := &chatgptgo.Request{
            Model: chatgptgo.Turbo,
            Messages: []*chatgptgo.Message{
                {
                    Role:    "user",
                    Content: "Hello, world!",
                },
            },
        }

        response, err := api.Chat(request)
        if err != nil {
            panic(err)
        }

        println(response.Choices[0].Message.Content)
    }
```

## Additional Configuration

Following configuration options are available in the `Request` struct:
- `Model`: ID of the model to use. Currently, only `gpt-3.5-turbo` and `gpt-3.5-turbo-0301` are supported.
- `Messages`: messages to generate chat completions for, in the chat format.
- `Temperature`: what sampling temperature to use, between 0 and 2. Higher values like 0.8 will make the output more random, while lower values like 0.2 will make it more focused and deterministic.
- `TopP`: an alternative to sampling with temperature, called nucleus sampling, where the model considers the results of the tokens with top_p probability mass.
- `N`: how many chat completion choices to generate for each input message.
- `Stop`: up to 4 sequences where the API will stop generating further tokens.
- `MaxTokens`: maximum number of tokens to generate for each chat completion choice.
- `PresencePenalty`: a number between -2.0 and 2.0. Positive values penalize new tokens based on whether they appear in the text so far, increasing the model's likelihood to talk about new topics.
- `FrequencyPenalty`: a number between -2.0 and 2.0. Positive values penalize new tokens based on their existing frequency in the text so far, decreasing the model's likelihood to repeat the same line verbatim.
- `User`: a unique identifier representing your end-user, which can help OpenAI to monitor and detect abuse.



## Providing Organization ID

You may provide your OpenAI organization ID when creating the API instance.
```go
    api := chatgptgo.NewApi("YOUR_API_KEY").WithOrganizationId("YOUR_ORGANIZATION_ID")
```


## Using your own HTTP client

You may provide your own custom HTTP client when creating the API instance. (Uses `http.DefaultClient` if not provided).
```go
    api := chatgptgo.NewApi("YOUR_API_KEY").WithClient(&http.Client{Timeout: 10 * time.Second})
```
