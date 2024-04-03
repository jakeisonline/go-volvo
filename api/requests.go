package api

import (
	"fmt"
	"net/http"
)

func Call(endpointName string, vin string) (*http.Response, error) {
	var resp *http.Response
	var err error
	endpoint := GetEndpoint(endpointName, vin)

	if endpoint.Method == "GET" {
		resp, err = get(endpoint.Url)
	} else {
		errorString := fmt.Sprintf("[go-volvo] Invalid method: %v", endpoint.Method)
		panic(errorString)
	}

	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusOK {
		return resp, err
	} else {
		return nil, fmt.Errorf("[go-volvo] status code received: %v", resp.StatusCode)
	}
}

func get(endpointUrl string) (*http.Response, error) {
	resp, err := http.Get(endpointUrl)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
