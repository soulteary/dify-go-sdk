package dify

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type DifyClientConfig struct {
	Key     string
	Host    string
	Timeout int
	SkipTLS bool
	User    string
}

type DifyClient struct {
	Key     string
	Host    string
	Timeout time.Duration
	SkipTLS bool
	Client  *http.Client
	User    string
}

func CreateDifyClient(config DifyClientConfig) (*DifyClient, error) {
	key := strings.TrimSpace(config.Key)
	if key == "" {
		return nil, fmt.Errorf("dify API Key is required")
	}

	host := strings.TrimSpace(config.Host)
	if host == "" {
		return nil, fmt.Errorf("dify Host is required")
	}

	timeout := 0 * time.Second
	if config.Timeout <= 0 {
		if config.Timeout < 0 {
			fmt.Println("Timeout should be a positive number, reset to default value: 10s")
		}
		timeout = DEFAULT_TIMEOUT * time.Second
	}

	skipTLS := false
	if config.SkipTLS {
		skipTLS = true
	}

	config.User = strings.TrimSpace(config.User)
	if config.User == "" {
		config.User = DEFAULT_USER
	}

	var client *http.Client

	if skipTLS {
		client = &http.Client{Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}}
	} else {
		client = &http.Client{}
	}

	if timeout > 0 {
		client.Timeout = timeout
	}

	return &DifyClient{
		Key:     key,
		Host:    host,
		Timeout: timeout,
		SkipTLS: skipTLS,
		Client:  client,
		User:    config.User,
	}, nil
}
