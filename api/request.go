package api

import (
	"net/http"
)

func request(endpointName string, vin string, c *Client) (interface{}, error) {
	endpointMap := GetEndpoint(endpointName, vin)

	req, err := http.NewRequest(endpointMap.Method, endpointMap.Url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+c.accessToken)
	req.Header.Add("VCC-API-Key", c.apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return handleResponse(endpointName, resp)
}
