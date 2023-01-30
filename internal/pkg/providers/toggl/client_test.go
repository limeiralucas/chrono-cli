package toggl_test

import (
	"testing"

	"github.com/limeiralucas/chrono-cli/internal/pkg/providers/toggl"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	client toggl.Client
}

func Test_Client(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (s *TestSuite) SetupSuite() {
	s.client = toggl.NewClient("fake-api-key")
}

func (s *TestSuite) Test_Get() {
}
