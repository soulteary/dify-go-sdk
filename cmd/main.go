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
}

func CompletionMessages(client *dify.DifyClient) {
	payload, err := dify.PrepareCompletionPayload(map[string]interface{}{"query": "hey"})
	if err != nil {
		log.Fatalf("failed to prepare payload: %v\n", err)
		return
	}

	// normal response
	response, err := client.CompletionMessages(payload, dify.RESPONSE_MODE_STREAMING, "abc-123", nil)
	if err != nil {
		log.Fatalf("failed to get completion messages: %v\n", err)
		return
	}
	fmt.Println(response)
	fmt.Println()

	// streaming response
	response, err = client.CompletionMessagesStreaming(payload, dify.RESPONSE_MODE_BLOCKING, "abc-123", nil)
	if err != nil {
		log.Fatalf("failed to get completion messages: %v\n", err)
		return
	}
	fmt.Println(response)
	fmt.Println()
}
