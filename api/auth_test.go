package api

import (
	"testing"
)

func TestSetAccessToken(t *testing.T) {
	testAccessToken := "1234567890"
	SetAccessToken(testAccessToken)

	if accessToken != testAccessToken {
		t.Errorf("SetAccessToken(\"%s\") did not set accessToken to expected value", testAccessToken)
	}
}

func TestSetApiKey(t *testing.T) {
	testApiKey := "0987654321"
	SetApiKey(testApiKey)

	if apiKey != testApiKey {
		t.Errorf("SetApiKey(\"%s\") did not set apiKey to expected value", testApiKey)
	}
}

func TestGetApiKeyPanic(t *testing.T) {
	apiKey = ""
	defer func() { _ = recover() }()
	result := GetApiKey()

	t.Errorf("GetApiKey did not panic with empty apiKey. %s", result)
}

func TestGetAccessTokenPanic(t *testing.T) {
	accessToken = ""
	defer func() { _ = recover() }()
	result := GetAccessToken()

	t.Errorf("GetAccessToken did not panic with empty accessToken. %s", result)
}
