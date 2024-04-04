package api

import (
	"fmt"
	"net/http"

	"github.com/jakeisonline/go-volvo/helpers"
)

func Call(endpointName string, vin string) (*http.Response, error) {
	endpoint := GetEndpoint(endpointName, vin)

	if endpoint.Method == "GET" || endpoint.Method == "POST" {
		resp, err := request(endpoint)
		return resp, err
	} else {
		panic(fmt.Sprintf("Invalid method: %v", endpoint.Method))
	}
}

func request(endpoint Endpoint) (*http.Response, error) {
	req, err := http.NewRequest(endpoint.Method, endpoint.Url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+GetAccessToken())
	req.Header.Add("VCC-API-Key", GetApiKey())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusOK {
		return resp, err
	} else {
		return nil, helpers.Error(fmt.Sprintf("status code received: %v\n\n%v", resp.StatusCode, resp))
	}
}
