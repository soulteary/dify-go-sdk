package dify

// https://dify.lab.io/console/api/workspaces/current/models/model-types/rerank

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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

	req, err := http.NewRequest("GET", api, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", dc.ConsoleToken))
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
