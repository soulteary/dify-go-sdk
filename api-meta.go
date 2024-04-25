package dify

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type GetMetaResponse struct {
	ToolIcons struct {
		Dalle2  string `json:"dalle2"`
		APITool struct {
			Background string `json:"background"`
			Content    string `json:"content"`
		} `json:"api_tool"`
	} `json:"tool_icons"`
}

func (dc *DifyClient) GetMeta() (result GetMetaResponse, err error) {
	req, err := http.NewRequest("GET", dc.GetAPI(API_META), nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", dc.Key))
	req.Header.Set("Content-Type", "application/json")

	resp, err := dc.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

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
