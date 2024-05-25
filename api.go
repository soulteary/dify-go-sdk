package dify

import "strings"

const (
	API_COMPLETION_MESSAGES      = "/completion-messages"
	API_COMPLETION_MESSAGES_STOP = "/completion-messages/:task_id/stop"

	API_CHAT_MESSAGES      = "/chat-messages"
	API_CHAT_MESSAGES_STOP = "/chat-messages/:task_id/stop"

	API_MESSAGES           = "/messages"
	API_MESSAGES_SUGGESTED = "/messages/:message_id/suggested"
	API_MESSAGES_FEEDBACKS = "/messages/:message_id/feedbacks"

	API_CONVERSATIONS        = "/conversations"
	API_CONVERSATIONS_DELETE = "/conversations/:conversation_id"
	API_CONVERSATIONS_RENAME = "/conversations/:conversation_id/name"

	API_FILE_UPLOAD = "/files/upload"
	API_PARAMETERS  = "/parameters"
	API_META        = "/meta"

	API_AUDIO_TO_TEXT = "/audio-to-text"
	API_TEXT_TO_AUDIO = "/text-to-audio"

	API_PARAM_TASK_ID         = ":task_id"
	API_PARAM_MESSAGE_ID      = ":message_id"
	API_PARAM_CONVERSATION_ID = ":conversation_id"

	CONSOLE_API_FILE_UPLOAD = "/files/upload?source=datasets"
	CONSOLE_API_LOGIN       = "/login"

	CONSOLE_API_PARAM_DATASETS_ID = ":datasets_id"

	CONSOLE_API_DATASETS_CREATE      = "/datasets"
	CONSOLE_API_DATASETS_LIST        = "/datasets"
	CONSOLE_API_DATASETS_DELETE      = "/datasets/:datasets_id"
	CONSOLE_API_DATASETS_INIT        = "/datasets/init"
	CONSOLE_API_DATASETS_INIT_STATUS = "/datasets/:datasets_id/indexing-status"

	CONSOLE_API_WORKSPACES_RERANK_MODEL        = "/workspaces/current/models/model-types/rerank"
	CONSOLE_API_CURRENT_WORKSPACE_RERANK_MODEL = "/workspaces/current/default-model?model_type=rerank"
)

func (dc *DifyClient) GetAPI(api string) string {
	return dc.Host + api
}

func (dc *DifyClient) GetConsoleAPI(api string) string {
	return dc.ConsoleHost + api
}

func UpdateAPIParam(api, key, value string) string {
	return strings.ReplaceAll(api, key, value)
}
