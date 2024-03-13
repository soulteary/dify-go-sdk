package dify

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
)

func (dc *DifyClient) TextToAudio(text string) (result any, err error) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("text", text)
	_ = writer.WriteField("user", dc.User)
	_ = writer.WriteField("streaming", "false")
	err = writer.Close()
	if err != nil {
		return result, fmt.Errorf("error closing writer: %v", err)
	}

	req, err := http.NewRequest("POST", dc.GetAPI(API_TEXT_TO_AUDIO), payload)
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
	return body, nil
}

func (dc *DifyClient) TextToAudioStreaming(text string) (result any, err error) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("text", text)
	_ = writer.WriteField("user", dc.User)
	_ = writer.WriteField("streaming", "true")
	err = writer.Close()
	if err != nil {
		return result, fmt.Errorf("error closing writer: %v", err)
	}

	req, err := http.NewRequest("POST", dc.GetAPI(API_TEXT_TO_AUDIO), payload)
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
	return body, nil
}
