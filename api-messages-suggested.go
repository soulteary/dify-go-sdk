package dify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type MessagesSuggestedResponse struct {
	Result string   `json:"result"`
	Data   []string `json:"data"`
}

func (dc *DifyClient) MessagesSuggested(message_id string) (result MessagesSuggestedResponse, err error) {
	if message_id == "" {
		return result, fmt.Errorf("message_id is required")
	}

	payloadBody := map[string]string{
		"user": dc.User,
	}

	api := dc.GetAPI(API_MESSAGES_SUGGESTED)
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
