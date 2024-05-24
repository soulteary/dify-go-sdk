package dify

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func setConsoleAuthorization(dc *DifyClient, req *http.Request) {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", dc.ConsoleToken))
	req.Header.Set("Content-Type", "application/json")
}

func SendGetRequestToConsole(dc *DifyClient, api string) (httpCode int, bodyText []byte, err error) {
	req, err := http.NewRequest("GET", api, nil)
	if err != nil {
		return -1, nil, err
	}

	setConsoleAuthorization(dc, req)

	resp, err := dc.Client.Do(req)
	if err != nil {
		return -1, nil, err
	}
	defer resp.Body.Close()

	bodyText, err = io.ReadAll(resp.Body)
	return resp.StatusCode, bodyText, err
}

func SendPostRequestToConsole(dc *DifyClient, api string, postBody interface{}) (httpCode int, bodyText []byte, err error) {
	var payload *strings.Reader
	if postBody != nil {
		buf, err := json.Marshal(postBody)
		if err != nil {
			return -1, nil, err
		}
		payload = strings.NewReader(string(buf))
	} else {
		payload = nil
	}

	req, err := http.NewRequest("POST", api, payload)
	if err != nil {
		return -1, nil, err
	}

	setConsoleAuthorization(dc, req)

	resp, err := dc.Client.Do(req)
	if err != nil {
		return -1, nil, err
	}
	defer resp.Body.Close()

	bodyText, err = io.ReadAll(resp.Body)
	return resp.StatusCode, bodyText, err
}

func CommonRiskForSendRequest(code int, err error) error {
	if err != nil {
		return err
	}

	if code != http.StatusOK {
		return fmt.Errorf("status code: %d", code)
	}

	return nil
}

func CommonRiskForSendRequestWithCode(code int, err error, targetCode int) error {
	if err != nil {
		return err
	}

	if code != targetCode {
		return fmt.Errorf("status code: %d", code)
	}

	return nil
}
