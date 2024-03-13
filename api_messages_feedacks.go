package dify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type MessagesFeedbacksResponse struct {
	Result string `json:"result"`
}

func (dc *DifyClient) MessagesFeedbacks(message_id string, rating string) (result MessagesFeedbacksResponse, err error) {
	if message_id == "" {
		return result, fmt.Errorf("message_id is required")
	}

	payloadBody := map[string]string{
		"user":   dc.User,
		"rating": rating,
	}

	api := dc.GetAPI(API_MESSAGES_FEEDBACKS)
	api = UpdateAPIParam(api, API_PARAM_MESSAGE_ID, message_id)

	buf, err := json.Marshal(payloadBody)
	if err != nil {
		return result, err
	}
	req, err := http.NewRequest("POST", api, bytes.NewBuffer(buf))
	if err != nil {
		return result, fmt.Errorf("could not create a new request: %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", dc.Key))
	req.Header.Set("Content-Type", "application/json")

	resp, err := dc.Client.Do(req)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		bodyText, err := io.ReadAll(resp.Body)
		if err != nil {
			return result, fmt.Errorf("status code: %d, could not read the body", resp.StatusCode)
		}
		return result, fmt.Errorf("status code: %d, %s", resp.StatusCode, bodyText)
	}

	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(bodyText, &result)
	if err != nil {
		return result, fmt.Errorf("failed to unmarshal the response: %v", err)
	}
	return result, nil
}
