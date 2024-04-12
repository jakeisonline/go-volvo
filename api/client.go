package api

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

func (c *Client) Call(endpointName string, vin string) (interface{}, error) {
	resp, err := request(endpointName, vin, c)
	return resp, err
}
