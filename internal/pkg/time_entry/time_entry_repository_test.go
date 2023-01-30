package time_entry_test

import (
	"errors"
	"testing"
	"time"

	core "github.com/limeiralucas/chrono-cli/internal/core/tag"
	"github.com/limeiralucas/chrono-cli/internal/pkg/time_entry"
	"github.com/limeiralucas/chrono-cli/internal/pkg/time_entry/mocks"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	provider *mocks.TimeEntryProvider
	repo     time_entry.TimeEntryRepository
	now      time.Time
}

func (s *TestSuite) SetupSuite() {
	s.now = time.Now()
}

func (s *TestSuite) ResetMocks() {
	s.provider = mocks.NewTimeEntryProvider(s.T())
	s.repo = time_entry.NewTimeEntryRepository(s.provider)
}

func TestTimeEntryRepository(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (s *TestSuite) Test_Create() {
	t := s.T()

	type ExpectedReturn struct {
		expectedId  int
		expectedErr error
	}

	cases := map[string]ExpectedReturn{
		"should call provider.Create and return id": {
			expectedId:  1,
			expectedErr: nil,
		},
		"should forward error from provider.Create": {
			expectedId:  0,
			expectedErr: errors.New("error on provider.Create"),
		},
	}

	for name, er := range cases {
		s.ResetMocks()
		s.Run(name, func() {
			s.provider.On("Create", "entry", s.now).Return(er.expectedId, er.expectedErr).Once()

			id, err := s.repo.Create("entry", s.now)

			s.provider.AssertNumberOfCalls(t, "Create", 1)
			s.Equal(er.expectedId, id)
			s.Equal(err, er.expectedErr)
		})
	}
}

func (s *TestSuite) Test_Delete() {
	t := s.T()

	cases := map[string]error{
		"should call provider.Delete":               nil,
		"should forward error from provider.Delete": errors.New("error on provider.Delete"),
	}

	for name, expectedErr := range cases {
		s.ResetMocks()
		s.Run(name, func() {
			s.provider.On("Delete", 1).Return(expectedErr).Once()
			err := s.repo.Delete(1)

			s.provider.AssertNumberOfCalls(t, "Delete", 1)
			s.Equal(expectedErr, err)
		})
	}
}

func (s *TestSuite) Test_GetTags() {
	t := s.T()

	type ExpectedReturn struct {
		expectedTags []core.Tag
		expectedErr  error
	}

	cases := map[string]ExpectedReturn{
		"should call provider.GetTags and return tags": {
			expectedTags: []core.Tag{
				{Id: 1, Name: "Tag 1"},
				{Id: 2, Name: "Tag 2"},
			},
			expectedErr: nil,
		},
		"should forward error from provider.GetTags": {
			expectedTags: nil,
			expectedErr:  errors.New("error on provider.GetTags"),
		},
	}

	for name, er := range cases {
		s.ResetMocks()
		s.Run(name, func() {
			s.provider.On("GetTags", 1).Return(er.expectedTags, er.expectedErr).Once()

			tags, err := s.repo.GetTags(1)

			s.provider.AssertNumberOfCalls(t, "GetTags", 1)
			s.Equal(er.expectedTags, tags)
			s.Equal(err, er.expectedErr)
		})
	}
}

func (s *TestSuite) Test_AddTags() {
	t := s.T()

	cases := map[string]error{
		"should call provider.AddTags":               nil,
		"should forward error from provider.AddTags": errors.New("error on provider.AddTags"),
	}

	for name, expectedErr := range cases {
		s.ResetMocks()
		s.Run(name, func() {
			newTags := []core.Tag{
				{Id: 1, Name: "Tag 1"},
				{Id: 2, Name: "Tag 2"},
			}

			s.provider.On("AddTags", 1, newTags).Return(expectedErr).Once()
			err := s.repo.AddTags(1, newTags)

			s.provider.AssertNumberOfCalls(t, "AddTags", 1)
			s.Equal(expectedErr, err)
		})
	}
}

func (s *TestSuite) Test_RemoveTags() {
	t := s.T()

	cases := map[string]error{
		"should call provider.RemoveTags":               nil,
		"should forward error from provider.RemoveTags": errors.New("error on provider.RemoveTags"),
	}

	for name, expectedErr := range cases {
		s.ResetMocks()
		s.Run(name, func() {
			tagsToBeRemoved := []core.Tag{
				{Id: 1, Name: "Tag 1"},
				{Id: 2, Name: "Tag 2"},
			}

			s.provider.On("RemoveTags", 1, tagsToBeRemoved).Return(expectedErr).Once()
			err := s.repo.RemoveTags(1, tagsToBeRemoved)

			s.provider.AssertNumberOfCalls(t, "RemoveTags", 1)
			s.Equal(expectedErr, err)
		})
	}
}

func (s *TestSuite) Test_UpdateTitle() {
	t := s.T()

	cases := map[string]error{
		"should call provider.UpdateTitle":               nil,
		"should forward error from provider.UpdateTitle": errors.New("error on provider.UpdateTitle"),
	}

	for name, expectedErr := range cases {
		s.ResetMocks()
		s.Run(name, func() {
			s.provider.On("UpdateTitle", 1, "New Title").Return(expectedErr).Once()
			err := s.repo.UpdateTitle(1, "New Title")

			s.provider.AssertNumberOfCalls(t, "UpdateTitle", 1)
			s.Equal(expectedErr, err)
		})
	}
}
