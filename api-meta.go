package dify

import (
	"encoding/json"
	"fmt"
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
	api := dc.GetAPI(API_META)
	code, body, err := SendGetRequestToAPI(dc, api)

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
