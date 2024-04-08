package api

import (
	"net/http"
)

func Call(endpointName string, vin string) (interface{}, error) {
	resp, err := request(endpointName, vin)
	return resp, err
}

func request(endpointName string, vin string) (interface{}, error) {
	endpointMap := GetEndpoint(endpointName, vin)

	req, err := http.NewRequest(endpointMap.Method, endpointMap.Url, nil)
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
	defer resp.Body.Close()

	return handleResponse(endpointName, resp)
}
