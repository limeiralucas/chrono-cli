package toggl

import (
	"fmt"
	"io"
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

func (c *Client) Get(path string, query map[string]string) ([]byte, error) {
	url := fmt.Sprintf("%s%s", API_URL, path)

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	request.SetBasicAuth(c.ApiKey, "api_token")

	if query != nil {
		urlQuery := request.URL.Query()

		for name, value := range query {
			urlQuery.Add(name, value)
		}
		request.URL.RawQuery = urlQuery.Encode()
	}

	response, err := c.HttpClient.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("request to %s; status code: %d; body: %s", request.URL.String(), response.StatusCode, string(body))
	}

	return body, nil
}
