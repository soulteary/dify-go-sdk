package dify

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type UserLoginParams struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	RememberMe bool   `json:"remember_me"`
}

type UserLoginResponse struct {
	Result string `json:"result"`
	Data   string `json:"data"`
}

func (dc *DifyClient) UserLogin(email string, password string) (result UserLoginResponse, err error) {
	var payload = UserLoginParams{
		Email:      email,
		Password:   password,
		RememberMe: true,
	}

	buf, err := json.Marshal(payload)
	if err != nil {
		return result, err
	}

	req, err := http.NewRequest("POST", dc.GetConsoleAPI(CONSOLE_API_LOGIN), strings.NewReader(string(buf)))
	if err != nil {
		return result, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", ""))
	req.Header.Set("Content-Type", "application/json")

	resp, err := dc.Client.Do(req)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
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
