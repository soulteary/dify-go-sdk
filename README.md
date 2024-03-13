# Dify Go SDK

Golang SDK for [langgenius/dify](https://github.com/langgenius/dify) .

## Quick Start

```bash
go get soulteary/dify-go-sdk
```

Create a Dify client that can invoke various capabilities.

```go
APIKey := "your_dify_api_key"
APIHost := "http://your-host/v1"

client, err := dify.CreateDifyClient(dify.DifyClientConfig{Key: APIKey, Host: APIHost})
if err != nil {
    log.Fatalf("failed to create DifyClient: %v\n", err)
    return
}
```

see [example](./cmd/main.go).

## Dify Client Config

```go
type DifyClientConfig struct {
	Key     string // API Key
	Host    string // API Host
	Timeout int    // Client Request Timeout
	SkipTLS bool   // Skip TLS Certs Verify (self-sign certs)
	User    string // AppId, for analytics
}
```

## API: `/completion-messages`

The most commonly used interface, used to call the model to generate content.

- CompletionMessages
- CompletionMessagesStreaming

```go
payload, err := dify.PrepareCompletionPayload(map[string]interface{}{"query": "hey"})
if err != nil {
    log.Fatalf("failed to prepare payload: %v\n", err)
    return
}

// normal response
completionMessagesResponse, err := client.CompletionMessages(payload, "", nil)
if err != nil {
    log.Fatalf("failed to get completion messages: %v\n", err)
    return
}
fmt.Println(completionMessagesResponse)
fmt.Println()

// streaming response
completionMessagesStreamingResponse, err := client.CompletionMessagesStreaming(payload, "", nil)
if err != nil {
    log.Fatalf("failed to get completion messages: %v\n", err)
    return
}
fmt.Println(completionMessagesStreamingResponse)
fmt.Println()
```

## API: `/files/upload`

Upload files to Dify.

- FileUpload

```go
fileUploadResponse, err := client.FileUpload("./README.md", "readme.md")
if err != nil {
    log.Fatalf("failed to upload file: %v\n", err)
    return
}
fmt.Println(fileUploadResponse)
fmt.Println()
```

## API: `/completion-messages/:task_id/stop`

Interface for interrupting streaming output.

- CompletionMessagesStop

```go
completionMessagesStopResponse, err := client.CompletionMessagesStop("0d2bd315-d4de-476f-ad5e-faaa00d571ea")
if err != nil {
    log.Fatalf("failed to stop completion messages: %v\n", err)
    return
}
fmt.Println(completionMessagesStopResponse)
fmt.Println()
```

## API: `/messages/:message_id/feedbacks`

Perform f on the interface output results.

- MessagesFeedbacks

```go
messagesFeedbacksResponse, err := client.MessagesFeedbacks(messageID, "like")
if err != nil {
    log.Fatalf("failed to get messages feedbacks: %v\n", err)
    return
}
fmt.Println(messagesFeedbacksResponse)
fmt.Println()
```

## API: `/parameters`

Get Dify parameters.

- GetParameters

```go
parametersResponse, err := client.GetParameters()
if err != nil {
    log.Fatalf("failed to get parameters: %v\n", err)
    return
}
fmt.Println(parametersResponse)
fmt.Println()
```

## API: `/text-to-audio`

text to audio.

- TextToAudio
- TextToAudioStreaming

```go
textToAudioResponse, err := client.TextToAudio("hello world")
if err != nil {
    log.Fatalf("failed to get text to audio: %v\n", err)
    return
}
fmt.Println(textToAudioResponse)
fmt.Println()

textToAudioStreamingResponse, err := client.TextToAudioStreaming("hello world")
if err != nil {
    log.Fatalf("failed to get text to audio streaming: %v\n", err)
    return
}
fmt.Println(textToAudioStreamingResponse)
fmt.Println()
```
