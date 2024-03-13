package dify

const (
	API_COMPLETION_MESSAGES = "/completion-messages"
	API_FILE_UPLOAD         = "/files/upload"
)

func (dc *DifyClient) GetAPI(api string) string {
	return dc.Host + api
}
