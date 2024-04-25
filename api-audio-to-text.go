package dify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

type AudioToTextResponse struct {
	Text string `json:"text"`
}

func (dc *DifyClient) AudioToText(filePath string) (result AudioToTextResponse, err error) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)

	fw, err := writer.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		return result, fmt.Errorf("error creating form file: %v", err)
	}

	fd, err := os.Open(filePath)
	if err != nil {
		return result, fmt.Errorf("error opening file: %v", err)
	}
	defer fd.Close()

	_, err = io.Copy(fw, fd)
	if err != nil {
		return result, fmt.Errorf("error copying file: %v", err)
	}

	err = writer.Close()
	if err != nil {
		return result, fmt.Errorf("error closing writer: %v", err)
	}

	req, err := http.NewRequest("POST", dc.GetAPI(API_AUDIO_TO_TEXT), payload)
	if err != nil {
		return result, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", dc.Key))
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := dc.Client.Do(req)
	if err != nil {
		return result, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return result, fmt.Errorf("status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return result, fmt.Errorf("could not read the body: %v", err)
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, fmt.Errorf("failed to unmarshal the response: %v", err)
	}

	return result, nil
}
