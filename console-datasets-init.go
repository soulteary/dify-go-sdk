package dify

import (
	"encoding/json"
	"fmt"
)

type InitDatasetsPayload struct {
	DataSource        InitDatasetsPayloadDataSource     `json:"data_source"`
	IndexingTechnique string                            `json:"indexing_technique"`
	ProcessRule       InitDatasetsPayloadProcessRule    `json:"process_rule"`
	DocForm           string                            `json:"doc_form"`
	DocLanguage       string                            `json:"doc_language"`
	RetrievalModel    InitDatasetsPayloadRetrievalModel `json:"retrieval_model"`
}

type InitDatasetsPayloadDataSource struct {
	Type     string                                `json:"type"`
	InfoList InitDatasetsPayloadDataSourceInfoList `json:"info_list"`
}

type InitDatasetsPayloadDataSourceInfoList struct {
	DataSourceType string                                    `json:"data_source_type"`
	FileInfoList   InitDatasetsPayloadDataSourceFileInfoList `json:"file_info_list"`
}

type InitDatasetsPayloadDataSourceFileInfoList struct {
	FileIds []string `json:"file_ids"`
}

type InitDatasetsResponse struct {
	Dataset   InitDatasetsResponseDataset           `json:"dataset"`
	Documents []InitDatasetsResponseDatasetDocument `json:"documents"`
	Batch     string                                `json:"batch"`
}

type InitDatasetsPayloadProcessRule struct {
	Rules struct {
	} `json:"rules"`
	Mode string `json:"mode"`
}

type InitDatasetsPayloadRetrievalModel struct {
	SearchMethod          string                            `json:"search_method"`
	RerankingEnable       bool                              `json:"reranking_enable"`
	RerankingModel        InitDatasetsPayloadRerankingModel `json:"reranking_model"`
	TopK                  int                               `json:"top_k"`
	ScoreThresholdEnabled bool                              `json:"score_threshold_enabled"`
	ScoreThreshold        float64                           `json:"score_threshold"`
}

type InitDatasetsPayloadRerankingModel struct {
	RerankingProviderName string `json:"reranking_provider_name"`
	RerankingModelName    string `json:"reranking_model_name"`
}

type InitDatasetsResponseDataset struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	Description       string `json:"description"`
	Permission        string `json:"permission"`
	DataSourceType    string `json:"data_source_type"`
	IndexingTechnique string `json:"indexing_technique"`
	CreatedBy         string `json:"created_by"`
	CreatedAt         int    `json:"created_at"`
}

type InitDatasetsResponseDatasetDocument struct {
	ID             string `json:"id"`
	Position       int    `json:"position"`
	DataSourceType string `json:"data_source_type"`
	DataSourceInfo struct {
		UploadFileID string `json:"upload_file_id"`
	} `json:"data_source_info"`
	DatasetProcessRuleID string `json:"dataset_process_rule_id"`
	Name                 string `json:"name"`
	CreatedFrom          string `json:"created_from"`
	CreatedBy            string `json:"created_by"`
	CreatedAt            int    `json:"created_at"`
	Tokens               int    `json:"tokens"`
	IndexingStatus       string `json:"indexing_status"`
	Error                any    `json:"error"`
	Enabled              bool   `json:"enabled"`
	DisabledAt           any    `json:"disabled_at"`
	DisabledBy           any    `json:"disabled_by"`
	Archived             bool   `json:"archived"`
	DisplayStatus        string `json:"display_status"`
	WordCount            int    `json:"word_count"`
	HitCount             int    `json:"hit_count"`
	DocForm              string `json:"doc_form"`
}

func (dc *DifyClient) InitDatasetsByUploadFile(datasets_ids []string) (result InitDatasetsResponse, err error) {
	payload := &InitDatasetsPayload{
		DocForm:           "text_model",
		DocLanguage:       "Chinese",
		IndexingTechnique: "high_quality",
		ProcessRule: InitDatasetsPayloadProcessRule{
			Mode: "automatic",
		},
		RetrievalModel: InitDatasetsPayloadRetrievalModel{
			RerankingEnable: false,
			RerankingModel: InitDatasetsPayloadRerankingModel{
				RerankingProviderName: "",
				RerankingModelName:    "",
			},
			ScoreThresholdEnabled: false,
			ScoreThreshold:        0.5,
			SearchMethod:          "semantic_search",
			TopK:                  3,
		},
		DataSource: InitDatasetsPayloadDataSource{
			Type: "upload_file",
			InfoList: InitDatasetsPayloadDataSourceInfoList{
				DataSourceType: "upload_file",
				FileInfoList: InitDatasetsPayloadDataSourceFileInfoList{
					FileIds: datasets_ids,
				},
			},
		},
	}

	api := dc.GetConsoleAPI(CONSOLE_API_DATASETS_INIT)

	code, body, err := SendPostRequestToConsole(dc, api, payload)

	err = CommonRiskForSendRequest(code, err)
	if err != nil {
		fmt.Println("error: ", string(body))
		return result, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, fmt.Errorf("failed to unmarshal the response: %v", err)
	}
	return result, nil
}

type InitDatasetsIndexingStatusResponse struct {
	Data []InitDatasetsIndexingStatusData `json:"data"`
}

type InitDatasetsIndexingStatusData struct {
	ID                   string `json:"id"`
	IndexingStatus       string `json:"indexing_status"`
	ProcessingStartedAt  int    `json:"processing_started_at"`
	ParsingCompletedAt   any    `json:"parsing_completed_at"`
	CleaningCompletedAt  any    `json:"cleaning_completed_at"`
	SplittingCompletedAt any    `json:"splitting_completed_at"`
	CompletedAt          any    `json:"completed_at"`
	PausedAt             any    `json:"paused_at"`
	Error                any    `json:"error"`
	StoppedAt            any    `json:"stopped_at"`
	CompletedSegments    int    `json:"completed_segments"`
	TotalSegments        int    `json:"total_segments"`
}

func (dc *DifyClient) InitDatasetsIndexingStatus(datasets_id string) (result InitDatasetsIndexingStatusResponse, err error) {
	api := dc.GetConsoleAPI(CONSOLE_API_DATASETS_INIT_STATUS)
	api = UpdateAPIParam(api, CONSOLE_API_PARAM_DATASETS_ID, datasets_id)

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
