package dify

import (
	"encoding/json"
	"fmt"
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

	api := dc.GetConsoleAPI(CONSOLE_API_LOGIN)

	code, body, err := SendPostRequestToConsole(dc, api, payload)

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
