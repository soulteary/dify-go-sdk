package main

import (
	"fmt"
	"log"
	"os"

	dify "github.com/soulteary/dify-go-sdk"
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

	ConsoleHost := os.Getenv("DIFY_CONSOLE_HOST")

	client, err := dify.CreateDifyClient(dify.DifyClientConfig{Key: APIKey, Host: APIHost, ConsoleHost: ConsoleHost})
	if err != nil {
		log.Fatalf("failed to create DifyClient: %v\n", err)
		return
	}

	// msgID := CompletionMessages(client)
	// FileUpload(client)
	// CompletionMessagesStop(client)
	// MessagesFeedbacks(client, msgID)
	// GetParameters(client)
	// TextToAudio(client)

	CONSOLE_USER := os.Getenv("DIFY_CONSOLE_USER")
	CONSOLE_PASS := os.Getenv("DIFY_CONSOLE_PASS")
	if CONSOLE_USER != "" && CONSOLE_PASS != "" {
		log.Println("Get Console Token")
		token := GetUserToken(client, CONSOLE_USER, CONSOLE_PASS)
		if token == "" {
			log.Fatalf("failed to get console token\n")
		}
		client.ConsoleToken = token

		// Create datasets
		var datasetsID string
		log.Println("Create datasets")
		createResult, err := client.CreateDatasets("test datasets")
		if err != nil {
			log.Fatalf("failed to create datasets: %v\n", err)
			return
		}
		datasetsID = createResult.ID
		log.Println(createResult)

		// List datasets
		log.Println("List datasets")
		ListResult, err := client.ListDatasets(1, 30)
		if err != nil {
			log.Fatalf("failed to list datasets: %v\n", err)
			return
		}
		if len(ListResult.Data) == 0 {
			log.Fatalf("no datasets found\n")
			return
		}
		for _, dataset := range ListResult.Data {
			if dataset.ID == datasetsID {
				// Delete datasets
				log.Println("Delete datasets")
				result, err := client.DeleteDatasets(datasetsID)
				if err != nil {
					log.Fatalf("failed to delete datasets: %v\n", err)
					return
				}
				log.Println(result)
			}
		}

		// Get the list of rerank models
		log.Println("List rerank models")
		reRankModels, err := client.ListWorkspacesRerankModels()
		if err != nil {
			log.Println("failed to list rerank models:", err)
		} else {
			log.Println(reRankModels)
		}

		// UploadFileToDatasets(client)
	}
}

func CompletionMessages(client *dify.DifyClient) (messageID string) {
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

	return completionMessagesResponse.MessageID
}

func FileUpload(client *dify.DifyClient) {
	fileUploadResponse, err := client.FileUpload("./README.md", "readme.md")
	if err != nil {
		log.Fatalf("failed to upload file: %v\n", err)
		return
	}
	fmt.Println(fileUploadResponse)
	fmt.Println()
}

func CompletionMessagesStop(client *dify.DifyClient) {
	completionMessagesStopResponse, err := client.CompletionMessagesStop("0d2bd315-d4de-476f-ad5e-faaa00d571ea")
	if err != nil {
		log.Fatalf("failed to stop completion messages: %v\n", err)
		return
	}
	fmt.Println(completionMessagesStopResponse)
	fmt.Println()
}

func MessagesFeedbacks(client *dify.DifyClient, messageID string) {
	messagesFeedbacksResponse, err := client.MessagesFeedbacks(messageID, "like")
	if err != nil {
		log.Fatalf("failed to get messages feedbacks: %v\n", err)
		return
	}
	fmt.Println(messagesFeedbacksResponse)
	fmt.Println()
}

func GetParameters(client *dify.DifyClient) {
	parametersResponse, err := client.GetParameters()
	if err != nil {
		log.Fatalf("failed to get parameters: %v\n", err)
		return
	}
	fmt.Println(parametersResponse)
	fmt.Println()
}

func TextToAudio(client *dify.DifyClient) {
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
}

func GetUserToken(client *dify.DifyClient, email, password string) string {
	result, err := client.UserLogin(email, password)
	if err != nil {
		log.Fatalf("failed to login: %v\n", err)
		return ""
	}
	return result.Data
}

func UploadFileToDatasets(client *dify.DifyClient) {
	err := os.WriteFile("testfile-for-dify-database.txt", []byte("test file for dify database"), 0644)
	if err != nil {
		log.Fatalf("failed to create file: %v\n", err)
		return
	}
	result, err := client.DatasetsFileUpload("testfile-for-dify-database.txt", "testfile-for-dify-database.txt")
	if err != nil {
		log.Fatalf("failed to upload file to datasets: %v\n", err)
		return
	}
	fmt.Println(result)
}
