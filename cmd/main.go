package main

import (
	"fmt"
	"log"
	"os"
	dify "soulteary/dify-go-sdk"
)

func main() {
	APIKey := os.Getenv("DIFY_API_KEY")
	if APIKey == "" {
		fmt.Println("DIFY_API_KEY is required")
		return
	}

	APIHost := os.Getenv("DIFY_API_HOST")
	if APIHost == "" {
		fmt.Println("DIFY_API_HOST is required")
		return
	}

	client, err := dify.CreateDifyClient(dify.DifyClientConfig{Key: APIKey, Host: APIHost})
	if err != nil {
		log.Fatalf("failed to create DifyClient: %v\n", err)
		return
	}

	CompletionMessages(client)
	FileUpload(client)
	CompletionMessagesStop(client)
}

func CompletionMessages(client *dify.DifyClient) {
	payload, err := dify.PrepareCompletionPayload(map[string]interface{}{"query": "hey"})
	if err != nil {
		log.Fatalf("failed to prepare payload: %v\n", err)
		return
	}

	// normal response
	completionMessagesResponse, err := client.CompletionMessages(payload, "abc-123", "", nil)
	if err != nil {
		log.Fatalf("failed to get completion messages: %v\n", err)
		return
	}
	fmt.Println(completionMessagesResponse)
	fmt.Println()

	// streaming response
	completionMessagesStreamingResponse, err := client.CompletionMessagesStreaming(payload, "abc-123", "", nil)
	if err != nil {
		log.Fatalf("failed to get completion messages: %v\n", err)
		return
	}
	fmt.Println(completionMessagesStreamingResponse)
	fmt.Println()
}

func FileUpload(client *dify.DifyClient) {
	fileUploadResponse, err := client.FileUpload("./README.md", "readme.md", "abc-123")
	if err != nil {
		log.Fatalf("failed to upload file: %v\n", err)
		return
	}
	fmt.Println(fileUploadResponse)
	fmt.Println()
}

func CompletionMessagesStop(client *dify.DifyClient) {
	completionMessagesStopResponse, err := client.CompletionMessagesStop("0d2bd315-d4de-476f-ad5e-faaa00d571ea", "abc-123")
	if err != nil {
		log.Fatalf("failed to stop completion messages: %v\n", err)
		return
	}
	fmt.Println(completionMessagesStopResponse)
	fmt.Println()
}
