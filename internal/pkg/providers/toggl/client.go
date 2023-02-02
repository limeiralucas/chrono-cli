package toggl

import (
	"fmt"
	"net/http"
)

const API_URL = "https://api.track.toggl.com/api/v9"

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	ApiKey     string
	HttpClient HTTPClient
}

func NewClient(apiKey string) Client {
	return Client{ApiKey: apiKey, HttpClient: &http.Client{}}
}

func (c *Client) Get(path string, query map[string]string) (any, error) {
	url := fmt.Sprintf("%s%s", API_URL, path)

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	request.SetBasicAuth("api_key", c.ApiKey)

	if query != nil {
		urlQuery := request.URL.Query()

		for name, value := range query {
			urlQuery.Add(name, value)
		}
		request.URL.RawQuery = urlQuery.Encode()
	}

	_, err = c.HttpClient.Do(request)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
