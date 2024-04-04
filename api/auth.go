package api

import (
	"github.com/jakeisonline/go-volvo/helpers"
)

var accessToken string
var apiKey string

func SetAccessToken(newAccessToken string) {
	accessToken = newAccessToken
}

func SetApiKey(newApiKey string) {
	apiKey = newApiKey
}

func GetAccessToken() string {
	if accessToken == "" {
		helpers.Panic("No accessToken defined, did you set one?")
	}

	return accessToken
}

func GetApiKey() string {
	if apiKey == "" {
		helpers.Panic("No apiKey defined, did you set one?")
	}

	return apiKey
}
