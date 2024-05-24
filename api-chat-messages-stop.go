package dify

import (
	"encoding/json"
	"fmt"
)

type ChatMessagesStopResponse struct {
	Result string `json:"result"`
}

func (dc *DifyClient) ChatMessagesStop(task_id string) (result ChatMessagesStopResponse, err error) {
	if task_id == "" {
		return result, fmt.Errorf("task_id is required")
	}

	payload := map[string]string{
		"user": dc.User,
	}

	api := dc.GetAPI(API_CHAT_MESSAGES_STOP)
	api = UpdateAPIParam(api, API_PARAM_TASK_ID, task_id)

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
