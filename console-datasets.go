package dify

import (
	"bytes"
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
	payloadBody := &CreateDatasetsPayload{
		Name: datasets_name,
	}

	api := dc.GetConsoleAPI(CONSOLE_API_DATASETS_CREATE)

	buf, err := json.Marshal(payloadBody)
	if err != nil {
		return result, err
	}
	req, err := http.NewRequest("POST", api, bytes.NewBuffer(buf))
	if err != nil {
		return result, fmt.Errorf("could not create a new request: %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", dc.ConsoleToken))
	req.Header.Set("Content-Type", "application/json")

	resp, err := dc.Client.Do(req)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		bodyText, err := io.ReadAll(resp.Body)
		if err != nil {
			return result, fmt.Errorf("status code: %d, could not read the body", resp.StatusCode)
		}
		return result, fmt.Errorf("status code: %d, %s", resp.StatusCode, bodyText)
	}

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
