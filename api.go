package dify

const (
	API_COMPLETION_MESSAGES = "/completion-messages"
)

func (dc *DifyClient) GetAPI(api string) string {
	return dc.Host + api
}
