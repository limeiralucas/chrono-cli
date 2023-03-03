package toggl

import (
	"bytes"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/limeiralucas/chrono-cli/mocks"
	"github.com/limeiralucas/chrono-cli/pkg/domain"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TimeEntryTestSuite struct {
	suite.Suite
	provider TimeEntryProvider
}

func Test_TimeEntry(t *testing.T) {
	suite.Run(t, new(TimeEntryTestSuite))
}

func (s *TimeEntryTestSuite) SetupSuite() {
	Client = mocks.NewHTTPClient(s.T())
	s.provider = NewTimeEntryProvider("fake-api-key")
}

func createBody(jsonStr string) io.ReadCloser {
	return io.NopCloser(bytes.NewReader([]byte(jsonStr)))
}

func (s *TimeEntryTestSuite) Test_Get() {
	type TestCase struct {
		responseBody string
		statusCode   int
		expectedErr  error
		expectedData any
	}

	testCases := map[string]TestCase{
		"should get current time entry (running)": {
			responseBody: `{"id": 1, "description": "desc", "start": "2023-01-01T01:00:00+00:00"}`,
			expectedData: &domain.TimeEntry{Id: 1, Description: "desc", StartDate: domain.TimeEntryTime{Time: time.Date(2023, 01, 01, 01, 00, 00, 00, time.UTC)}},
			statusCode:   200,
			expectedErr:  nil,
		},
		// "should get current time entry (not running)": {
		// 	responseBody: "{}",
		// 	expectedData: nil,
		// 	statusCode:   200,
		// },
	}

	for name, c := range testCases {
		s.Run(name, func() {
			var req *http.Request
			mockResp := &http.Response{
				StatusCode: c.statusCode,
				Body:       createBody(c.responseBody),
			}
			Client.(*mocks.HTTPClient).On("Do", mock.AnythingOfType("*http.Request")).Run(func(args mock.Arguments) {
				req = args.Get(0).(*http.Request)
			}).Return(mockResp, nil).Once()

			data, err := s.provider.GetCurrent()
			username, password, _ := req.BasicAuth()

			s.Equal(c.expectedErr, err)
			s.Equal("fake-api-key", username)
			s.Equal("api_token", password)
			s.Equal(c.expectedData, data)
		})
	}
}
