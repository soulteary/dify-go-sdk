package dify

import "strings"

const (
	API_COMPLETION_MESSAGES      = "/completion-messages"
	API_COMPLETION_MESSAGES_STOP = "/completion-messages/:task_id/stop"

	API_CHAT_MESSAGES      = "/chat-messages"
	API_CHAT_MESSAGES_STOP = "/chat-messages/:task_id/stop"

	API_MESSAGES_SUGGESTED = "/messages/:message_id/suggested"
	API_MESSAGES_FEEDBACKS = "/messages/:message_id/feedbacks"

	API_FILE_UPLOAD   = "/files/upload"
	API_PARAMETERS    = "/parameters"
	API_TEXT_TO_AUDIO = "/text-to-audio"

	API_PARAM_TASK_ID    = ":task_id"
	API_PARAM_MESSAGE_ID = ":message_id"
)

func (dc *DifyClient) GetAPI(api string) string {
	return dc.Host + api
}

func UpdateAPIParam(api, key, value string) string {
	return strings.ReplaceAll(api, key, value)
}
