package dify

// https://dify.lab.io/console/api/workspaces/current/models/model-types/rerank

import (
	"encoding/json"
	"fmt"
)

type ListWorkspacesRerankModelsResponse struct {
	Data []ListWorkspacesRerankItem `json:"data"`
}

type ListWorkspacesRerankItem struct {
	Provider string `json:"provider"`
	Label    struct {
		ZhHans string `json:"zh_Hans"`
		EnUS   string `json:"en_US"`
	} `json:"label"`
	IconSmall struct {
		ZhHans string `json:"zh_Hans"`
		EnUS   string `json:"en_US"`
	} `json:"icon_small"`
	IconLarge struct {
		ZhHans string `json:"zh_Hans"`
		EnUS   string `json:"en_US"`
	} `json:"icon_large"`
	Status string                      `json:"status"`
	Models []ListWorkspacesRerankModel `json:"models"`
}

type ListWorkspacesRerankModel struct {
	Model string `json:"model"`
	Label struct {
		ZhHans string `json:"zh_Hans"`
		EnUS   string `json:"en_US"`
	} `json:"label"`
	ModelType       string `json:"model_type"`
	Features        any    `json:"features"`
	FetchFrom       string `json:"fetch_from"`
	ModelProperties struct {
		ContextSize int `json:"context_size"`
	} `json:"model_properties"`
	Deprecated bool   `json:"deprecated"`
	Status     string `json:"status"`
}

func (dc *DifyClient) ListWorkspacesRerankModels() (result ListWorkspacesRerankModelsResponse, err error) {
	api := dc.GetConsoleAPI(CONSOLE_API_WORKSPACES_RERANK_MODEL)

	code, body, err := SendGetRequestToConsole(dc, api)

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

type GetCurrentWorkspaceRerankDefaultModelResponse struct {
	Data any `json:"data"`
}

func (dc *DifyClient) GetCurrentWorkspaceRerankDefaultModel() (result GetCurrentWorkspaceRerankDefaultModelResponse, err error) {
	api := dc.GetConsoleAPI(CONSOLE_API_CURRENT_WORKSPACE_RERANK_MODEL)

	code, body, err := SendGetRequestToConsole(dc, api)

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
