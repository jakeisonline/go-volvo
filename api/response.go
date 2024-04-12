package api

import (
	"encoding/json"
)

func UnmarshalAndCheck(body []byte, v interface{}) (interface{}, error) {
	err := json.Unmarshal(body, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}
