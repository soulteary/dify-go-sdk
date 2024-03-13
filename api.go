package dify

import "strings"

const (
	API_COMPLETION_MESSAGES      = "/completion-messages"
	API_FILE_UPLOAD              = "/files/upload"
	API_COMPLETION_MESSAGES_STOP = "/completion-messages/:task_id/stop"

	API_PARAM_TASK_ID = ":task_id"
)

func (dc *DifyClient) GetAPI(api string) string {
	return dc.Host + api
}

func UpdateAPIParam(api, key, value string) string {
	return strings.ReplaceAll(api, key, value)
}
