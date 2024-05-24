package dify

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (dc *DifyClient) DeleteDatasets(datasets_id string) (ok bool, err error) {
	if datasets_id == "" {
		return false, fmt.Errorf("datasets_id is required")
	}

	api := dc.GetConsoleAPI(CONSOLE_API_DATASETS_DELETE)
	api = UpdateAPIParam(api, CONSOLE_API_PARAM_DATASETS_ID, datasets_id)

	req, err := http.NewRequest("DELETE", api, nil)
	if err != nil {
		return false, fmt.Errorf("could not create a new request: %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", dc.ConsoleToken))
	req.Header.Set("Content-Type", "application/json")

	resp, err := dc.Client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		bodyText, err := io.ReadAll(resp.Body)
		if err != nil {
			return false, fmt.Errorf("status code: %d, could not read the body", resp.StatusCode)
		}
		return false, fmt.Errorf("status code: %d, %s", resp.StatusCode, bodyText)
	}

	return true, nil
}

type CreateDatasetsPayload struct {
	Name string `json:"name"`
}

type CreateDatasetsResponse struct {
	ID                     string `json:"id"`
	Name                   string `json:"name"`
	Description            any    `json:"description"`
	Provider               string `json:"provider"`
	Permission             string `json:"permission"`
	DataSourceType         any    `json:"data_source_type"`
	IndexingTechnique      any    `json:"indexing_technique"`
	AppCount               int    `json:"app_count"`
	DocumentCount          int    `json:"document_count"`
	WordCount              int    `json:"word_count"`
	CreatedBy              string `json:"created_by"`
	CreatedAt              int    `json:"created_at"`
	UpdatedBy              string `json:"updated_by"`
	UpdatedAt              int    `json:"updated_at"`
	EmbeddingModel         any    `json:"embedding_model"`
	EmbeddingModelProvider any    `json:"embedding_model_provider"`
	EmbeddingAvailable     any    `json:"embedding_available"`
	RetrievalModelDict     struct {
		SearchMethod    string `json:"search_method"`
		RerankingEnable bool   `json:"reranking_enable"`
		RerankingModel  struct {
			RerankingProviderName string `json:"reranking_provider_name"`
			RerankingModelName    string `json:"reranking_model_name"`
		} `json:"reranking_model"`
		TopK                  int  `json:"top_k"`
		ScoreThresholdEnabled bool `json:"score_threshold_enabled"`
		ScoreThreshold        any  `json:"score_threshold"`
	} `json:"retrieval_model_dict"`
	Tags []any `json:"tags"`
}

func (dc *DifyClient) CreateDatasets(datasets_name string) (result CreateDatasetsResponse, err error) {
	payload := &CreateDatasetsPayload{
		Name: datasets_name,
	}

	api := dc.GetConsoleAPI(CONSOLE_API_DATASETS_CREATE)

	code, body, err := SendPostRequestToConsole(dc, api, payload)

	err = CommonRiskForSendRequestWithCode(code, err, http.StatusCreated)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, fmt.Errorf("failed to unmarshal the response: %v", err)
	}
	return result, nil
}

type ListDatasetsResponse struct {
	Page    int                `json:"page"`
	Limit   int                `json:"limit"`
	Total   int                `json:"total"`
	HasMore bool               `json:"has_more"`
	Data    []ListDatasetsItem `json:"data"`
}

type ListDatasetsItem struct {
	ID             string                  `json:"id"`
	Name           string                  `json:"name"`
	Description    string                  `json:"description"`
	Mode           string                  `json:"mode"`
	Icon           string                  `json:"icon"`
	IconBackground string                  `json:"icon_background"`
	ModelConfig    ListDatasetsModelConfig `json:"model_config"`
	CreatedAt      int                     `json:"created_at"`
	Tags           []any                   `json:"tags"`
}

type ListDatasetsModelConfig struct {
	Model     ListDatasetsModelConfigDetail `json:"model"`
	PrePrompt string                        `json:"pre_prompt"`
}

type ListDatasetsModelConfigDetail struct {
	Provider         string `json:"provider"`
	Name             string `json:"name"`
	Mode             string `json:"mode"`
	CompletionParams struct {
	} `json:"completion_params"`
}

func (dc *DifyClient) ListDatasets(page int, limit int) (result ListDatasetsResponse, err error) {
	if page < 1 {
		return result, fmt.Errorf("page should be greater than 0")
	}
	if limit < 1 {
		return result, fmt.Errorf("limit should be greater than 0")
	}

	api := dc.GetConsoleAPI(CONSOLE_API_DATASETS_LIST)
	api = fmt.Sprintf("%s?page=%d&limit=%d", api, page, limit)

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
