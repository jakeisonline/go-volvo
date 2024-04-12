package api

import (
	"fmt"
	"net/http"
)

type Configuration struct {
	ApiKey      string
	AccessToken string
}

type Client struct {
	apiKey      string
	accessToken string
}

func NewClient(apiKey string) *Client {
	return newClientConfig(
		Configuration{
			ApiKey: apiKey,
		},
	)
}

func newClientConfig(config Configuration) *Client {
	return &Client{
		apiKey:      config.ApiKey,
		accessToken: config.AccessToken,
	}
}

func (c *Client) SetAccessToken(newAccessToken string) {
	c.accessToken = newAccessToken
}

func (c *Client) Call(endpointName string, vin string) (*http.Response, error) {
	resp, err := request(endpointName, vin, c)

	if resp.StatusCode == http.StatusOK {
		return resp, err
	} else {
		return nil, fmt.Errorf("status code received: %v\n\n%v", resp.StatusCode, resp)
	}
}
