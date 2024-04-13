package api

import (
	"crypto/rand"
	"fmt"
	"net/url"
	"strings"
)

type AuthRedirect struct {
	RedirectUrl string
	State       string
}

func createAuthRedirectUrl(redirectUrl string, c *Client) AuthRedirect {
	escapedRedirectUrl := url.QueryEscape(redirectUrl)
	state := createStateString()
	scope := createScopeString()
	authRedirectUrl := fmt.Sprintf("https://volvoid.eu.volvocars.com/as/authorization.oauth2?response_type=code&client_id=%s&redirect_uri=%s&scope=%s&state=%s", c.apiKey, escapedRedirectUrl, scope, state)

	return AuthRedirect{
		RedirectUrl: authRedirectUrl,
		State:       state,
	}
}

func createStateString() string {
	bytes := make([]byte, 5)
	if _, err := rand.Read(bytes); err != nil {
		panic(err)
	}
	state := fmt.Sprintf("%X", bytes)
	return state
}

func createScopeString() string {
	returnScope := strings.Join(scopes[:], " ")
	escapedReturnScope := url.QueryEscape(returnScope)
	return escapedReturnScope
}
