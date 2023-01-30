package toggl

type Client struct {
	apiKey string
}

func NewClient(apiKey string) Client {
	return Client{apiKey: apiKey}
}

// func (c *Client) Get(path string) map[string]interface{} {

// }
