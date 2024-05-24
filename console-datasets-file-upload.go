package dify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func (dc *DifyClient) DatasetsFileUpload(filePath string, fileName string) (result FileUploadResponse, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return result, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", fileName)
	if err != nil {
		return result, fmt.Errorf("error creating form file: %v", err)
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return result, fmt.Errorf("error copying file: %v", err)
	}

	_ = writer.WriteField("user", dc.User)
	err = writer.Close()
	if err != nil {
		return result, fmt.Errorf("error closing writer: %v", err)
	}

	req, err := http.NewRequest("POST", dc.GetConsoleAPI(CONSOLE_API_FILE_UPLOAD), body)
	if err != nil {
		return result, fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", dc.ConsoleToken))
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return result, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		return result, fmt.Errorf("status code: %d, create file failed", resp.StatusCode)
	}

	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return result, fmt.Errorf("could not read the body: %v", err)
	}

	err = json.Unmarshal(bodyText, &result)
	if err != nil {
		return result, fmt.Errorf("failed to unmarshal the response: %v", err)
	}
	return result, nil
}
