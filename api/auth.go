package api

var accessToken string
var apiKey string

func SetAccessToken(newAccessToken string) {
	accessToken = newAccessToken
}

func SetApiKey(newApiKey string) {
	apiKey = newApiKey
}
