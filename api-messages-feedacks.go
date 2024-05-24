package dify

import (
	"encoding/json"
	"fmt"
)

type MessagesFeedbacksResponse struct {
	Result string `json:"result"`
}

func (dc *DifyClient) MessagesFeedbacks(message_id string, rating string) (result MessagesFeedbacksResponse, err error) {
	if message_id == "" {
		return result, fmt.Errorf("message_id is required")
	}

	payload := map[string]string{
		"user":   dc.User,
		"rating": rating,
	}

	api := dc.GetAPI(API_MESSAGES_FEEDBACKS)
	api = UpdateAPIParam(api, API_PARAM_MESSAGE_ID, message_id)

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
