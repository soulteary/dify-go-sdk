package dify

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type GetParametersResponse struct {
	OpeningStatement              string `json:"opening_statement"`
	SuggestedQuestions            []any  `json:"suggested_questions"`
	SuggestedQuestionsAfterAnswer struct {
		Enabled bool `json:"enabled"`
	} `json:"suggested_questions_after_answer"`
	SpeechToText struct {
		Enabled bool `json:"enabled"`
	} `json:"speech_to_text"`
	TextToSpeech struct {
		Enabled  bool   `json:"enabled"`
		Voice    string `json:"voice"`
		Language string `json:"language"`
	} `json:"text_to_speech"`
	RetrieverResource struct {
		Enabled bool `json:"enabled"`
	} `json:"retriever_resource"`
	AnnotationReply struct {
		Enabled bool `json:"enabled"`
	} `json:"annotation_reply"`
	MoreLikeThis struct {
		Enabled bool `json:"enabled"`
	} `json:"more_like_this"`
	UserInputForm []struct {
		Paragraph struct {
			Label    string `json:"label"`
			Variable string `json:"variable"`
			Required bool   `json:"required"`
			Default  string `json:"default"`
		} `json:"paragraph"`
	} `json:"user_input_form"`
	SensitiveWordAvoidance struct {
		Enabled bool   `json:"enabled"`
		Type    string `json:"type"`
		Configs []any  `json:"configs"`
	} `json:"sensitive_word_avoidance"`
	FileUpload struct {
		Image struct {
			Enabled         bool     `json:"enabled"`
			NumberLimits    int      `json:"number_limits"`
			Detail          string   `json:"detail"`
			TransferMethods []string `json:"transfer_methods"`
		} `json:"image"`
	} `json:"file_upload"`
	SystemParameters struct {
		ImageFileSizeLimit string `json:"image_file_size_limit"`
	} `json:"system_parameters"`
}

func (dc *DifyClient) GetParameters() (result GetParametersResponse, err error) {
	req, err := http.NewRequest("GET", dc.GetAPI(API_PARAMETERS), nil)
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
