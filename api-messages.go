package dify

import (
	"encoding/json"
	"fmt"
)

type MessagesResponse struct {
	Limit   int  `json:"limit"`
	HasMore bool `json:"has_more"`
	Data    []struct {
		ID             string `json:"id"`
		ConversationID string `json:"conversation_id"`
		Inputs         struct {
			Name string `json:"name"`
		} `json:"inputs"`
		Query              string `json:"query"`
		Answer             string `json:"answer"`
		MessageFiles       []any  `json:"message_files"`
		Feedback           any    `json:"feedback"`
		RetrieverResources []struct {
			Position     int     `json:"position"`
			DatasetID    string  `json:"dataset_id"`
			DatasetName  string  `json:"dataset_name"`
			DocumentID   string  `json:"document_id"`
			DocumentName string  `json:"document_name"`
			SegmentID    string  `json:"segment_id"`
			Score        float64 `json:"score"`
			Content      string  `json:"content"`
		} `json:"retriever_resources"`
		AgentThoughts []any `json:"agent_thoughts"`
		CreatedAt     int   `json:"created_at"`
	} `json:"data"`
}

func (dc *DifyClient) Messages(conversation_id string) (result MessagesResponse, err error) {
	if conversation_id == "" {
		return result, fmt.Errorf("conversation_id is required")
	}

	payloadBody := map[string]string{
		"user":            dc.User,
		"conversation_id": conversation_id,
	}

	api := dc.GetAPI(API_MESSAGES)

	code, body, err := SendPostRequestToAPI(dc, api, payloadBody)

	err = CommonRiskForSendRequest(code, err)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, fmt.Errorf("failed to unmarshal the response: %v", err)
	}
	return result, nil
}
