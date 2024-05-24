package dify

import (
	"fmt"
	"net/http"
)

func setAPIAuthorization(dc *DifyClient, req *http.Request) {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", dc.Key))
	req.Header.Set("Content-Type", "application/json")
}

func SendGetRequestToAPI(dc *DifyClient, api string) (httpCode int, bodyText []byte, err error) {
	return SendGetRequest(false, dc, api)
}

func SendPostRequestToAPI(dc *DifyClient, api string, postBody interface{}) (httpCode int, bodyText []byte, err error) {
	return SendPostRequest(false, dc, api, postBody)
}
