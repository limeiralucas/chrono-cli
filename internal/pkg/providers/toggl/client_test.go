package toggl_test

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/limeiralucas/chrono-cli/internal/pkg/providers/mocks"
	"github.com/limeiralucas/chrono-cli/internal/pkg/providers/toggl"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type ClientTestSuite struct {
	suite.Suite
	client     toggl.Client
	httpClient *mocks.HTTPClient
}

func Test_Client(t *testing.T) {
	suite.Run(t, new(ClientTestSuite))
}

func (s *ClientTestSuite) SetupSuite() {
	s.httpClient = mocks.NewHTTPClient(s.T())
	s.client = toggl.NewClient("fake-api-key")
	s.client.HttpClient = s.httpClient
}

func createBody(jsonStr string) io.ReadCloser {
	return io.NopCloser(bytes.NewReader([]byte(jsonStr)))
}

func (s *ClientTestSuite) Test_Get() {
	type TestCase struct {
		statusCode int
		body       string
		query      map[string]string
	}

	testCases := map[string]TestCase{
		"should make a authenticated request": {
			statusCode: 200,
			body:       "[]",
			query:      nil,
		},
		"should make a request with the provided query params": {
			statusCode: 200,
			body:       "[]",
			query: map[string]string{
				"param1": "1",
				"param2": "2",
			},
		},
		"should return the response body": {
			statusCode: 200,
			body:       `{"id": 1}`,
			query:      nil,
		},
		"should return the http status code": {
			statusCode: 401,
			body:       "[]",
			query:      nil,
		},
	}

	for name, c := range testCases {
		s.Run(name, func() {
			var req *http.Request
			mockResp := &http.Response{
				StatusCode: c.statusCode,
				Body:       createBody(c.body),
			}
			s.httpClient.On("Do", mock.AnythingOfType("*http.Request")).Run(func(args mock.Arguments) {
				req = args.Get(0).(*http.Request)
			}).Return(mockResp, nil).Once()

			body, statusCode, _ := s.client.Get("/", c.query)
			username, password, _ := req.BasicAuth()
			querySent := req.URL.Query()

			s.Equal("fake-api-key", username)
			s.Equal("api_token", password)
			s.Equal([]byte(c.body), body)
			s.Equal(c.statusCode, statusCode)

			for key, value := range c.query {
				s.Equal(querySent.Get(key), value)
			}
		})
	}
}
