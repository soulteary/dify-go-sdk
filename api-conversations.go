package dify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ConversationsResponse struct {
	Limit   int  `json:"limit"`
	HasMore bool `json:"has_more"`
	Data    []struct {
		ID     string `json:"id"`
		Name   string `json:"name,omitempty"`
		Inputs struct {
			Book   string `json:"book"`
			MyName string `json:"myName"`
		} `json:"inputs,omitempty"`
		Status    string `json:"status,omitempty"`
		CreatedAt int    `json:"created_at,omitempty"`
	} `json:"data"`
}

func (dc *DifyClient) Conversations(last_id string, limit int) (result ConversationsResponse, err error) {
	payloadLimit := ""
	if limit <= 0 {
		limit = 20
	}
	payloadLimit = fmt.Sprintf("%d", limit)

	payloadBody := map[string]string{
		"user":    dc.User,
		"last_id": last_id,
		"limit":   payloadLimit,
	}

	api := dc.GetAPI(API_CONVERSATIONS)

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

type DeleteConversationsResponse struct {
	Result string `json:"result"`
}

func (dc *DifyClient) DeleteConversation(conversation_id string) (result DeleteConversationsResponse, err error) {
	if conversation_id == "" {
		return result, fmt.Errorf("conversation_id is required")
	}

	payloadBody := map[string]string{
		"user": dc.User,
	}

	api := dc.GetAPI(API_CONVERSATIONS_DELETE)
	api = UpdateAPIParam(api, API_PARAM_CONVERSATION_ID, conversation_id)

	buf, err := json.Marshal(payloadBody)
	if err != nil {
		return result, err
	}
	req, err := http.NewRequest("DELETE", api, bytes.NewBuffer(buf))
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

type RenameConversationsResponse struct {
	Result string `json:"result"`
}

func (dc *DifyClient) RenameConversation(conversation_id string) (result RenameConversationsResponse, err error) {
	if conversation_id == "" {
		return result, fmt.Errorf("conversation_id is required")
	}

	payload := map[string]string{
		"user": dc.User,
	}

	api := dc.GetAPI(API_CONVERSATIONS_RENAME)
	api = UpdateAPIParam(api, API_PARAM_CONVERSATION_ID, conversation_id)

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
