package dify

import (
	"encoding/json"
	"fmt"
)

type ChatMessagesPayload struct {
	Inputs         any                       `json:"inputs"`
	Query          string                    `json:"query"`
	ResponseMode   string                    `json:"response_mode"`
	ConversationID string                    `json:"conversation_id,omitempty"`
	User           string                    `json:"user"`
	Files          []ChatMessagesPayloadFile `json:"files,omitempty"`
}

type ChatMessagesPayloadFile struct {
	Type           string `json:"type"`
	TransferMethod string `json:"transfer_method"`
	URL            string `json:"url,omitempty"`
	UploadFileID   string `json:"upload_file_id,omitempty"`
}

type ChatMessagesResponse struct {
	Event          string `json:"event"`
	MessageID      string `json:"message_id"`
	ConversationID string `json:"conversation_id"`
	Mode           string `json:"mode"`
	Answer         string `json:"answer"`
	Metadata       any    `json:"metadata"`
	CreatedAt      int    `json:"created_at"`
}

func PrepareChatPayload(payload map[string]interface{}) (string, error) {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func (dc *DifyClient) ChatMessages(inputs string, conversation_id string, files []any) (result ChatMessagesResponse, err error) {
	var payload ChatMessagesPayload

	if len(inputs) == 0 {
		return result, fmt.Errorf("inputs is required")
	} else {
		var tryDecode map[string]interface{}
		err := json.Unmarshal([]byte(inputs), &tryDecode)
		if err != nil {
			return result, fmt.Errorf("inputs should be a valid JSON string")
		}
		payload.Inputs = tryDecode
	}

	payload.ResponseMode = RESPONSE_MODE_BLOCKING
	payload.User = dc.User

	if conversation_id != "" {
		payload.ConversationID = conversation_id
	}

	if len(files) > 0 {
		// TODO TBD
		return result, fmt.Errorf("files are not supported")
	}

	api := dc.GetAPI(API_CHAT_MESSAGES)

	code, body, err := SendPostRequestToAPI(dc, api, payload)

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

func (dc *DifyClient) ChatMessagesStreaming(inputs string, conversation_id string, files []any) (result string, err error) {
	var payload ChatMessagesPayload

	if len(inputs) == 0 {
		return "", fmt.Errorf("inputs is required")
	} else {
		var tryDecode map[string]interface{}
		err := json.Unmarshal([]byte(inputs), &tryDecode)
		if err != nil {
			return "", fmt.Errorf("inputs should be a valid JSON string")
		}
		payload.Inputs = tryDecode
	}

	payload.ResponseMode = RESPONSE_MODE_STREAMING
	payload.User = dc.User

	if conversation_id != "" {
		payload.ConversationID = conversation_id
	}

	if len(files) > 0 {
		// TODO TBD
		return "", fmt.Errorf("files are not supported")
	}

	api := dc.GetAPI(API_CHAT_MESSAGES)

	code, body, err := SendPostRequestToAPI(dc, api, payload)

	err = CommonRiskForSendRequest(code, err)
	if err != nil {
		return result, err
	}

	// if !strings.Contains(resp.Header.Get("Content-Type"), "text/event-stream") {
	// 	return "", fmt.Errorf("response is not a streaming response")
	// }

	return string(body), nil
}
