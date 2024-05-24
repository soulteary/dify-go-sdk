package dify

import (
	"fmt"
	"net/http"
)

func setConsoleAuthorization(dc *DifyClient, req *http.Request) {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", dc.ConsoleToken))
	req.Header.Set("Content-Type", "application/json")
}

func SendGetRequestToConsole(dc *DifyClient, api string) (httpCode int, bodyText []byte, err error) {
	return SendGetRequest(true, dc, api)
}

func SendPostRequestToConsole(dc *DifyClient, api string, postBody interface{}) (httpCode int, bodyText []byte, err error) {
	return SendPostRequest(true, dc, api, postBody)
}
