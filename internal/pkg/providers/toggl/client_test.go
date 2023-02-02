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

type TestSuite struct {
	suite.Suite
	client     toggl.Client
	httpClient *mocks.HTTPClient
}

func TestClient(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (s *TestSuite) SetupSuite() {
	s.httpClient = mocks.NewHTTPClient(s.T())
	s.client = toggl.NewClient("fake-api-key")
	s.client.HttpClient = s.httpClient
}

func createBody(jsonStr string) io.ReadCloser {
	return io.NopCloser(bytes.NewReader([]byte(jsonStr)))
}

func (s *TestSuite) Test_GetAuthenticated() {
	var req *http.Request
	mockResp := &http.Response{
		StatusCode: 200,
		Body:       createBody("[]"),
	}
	s.httpClient.On("Do", mock.AnythingOfType("*http.Request")).Run(func(args mock.Arguments) {
		req = args.Get(0).(*http.Request)
	}).Return(mockResp, nil).Once()

	s.client.Get("/", nil)
	username, password, _ := req.BasicAuth()

	s.Equal("api_key", username)
	s.Equal("fake-api-key", password)
}

func (s *TestSuite) Test_GetWithQuery() {
	var req *http.Request
	mockResp := &http.Response{
		StatusCode: 200,
		Body:       createBody("[]"),
	}

	s.httpClient.On("Do", mock.AnythingOfType("*http.Request")).Run(func(args mock.Arguments) {
		req = args.Get(0).(*http.Request)
	}).Return(mockResp, nil).Once()

	s.client.Get("/", map[string]string{
		"param1": "1",
		"param2": "2",
	})

	querySent := req.URL.Query()

	s.Equal(querySent.Get("param1"), "1")
	s.Equal(querySent.Get("param2"), "2")
}
