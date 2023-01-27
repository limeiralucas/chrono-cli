package tag_test

import (
	"errors"
	"testing"

	"github.com/limeiralucas/chrono-cli/internal/tag"
	"github.com/limeiralucas/chrono-cli/mocks"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	provider *mocks.TagProvider
	repo     tag.TagRepository
}

func TestTagRepository(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (s *TestSuite) ResetMocks() {
	s.provider = mocks.NewTagProvider(s.T())
	s.repo = tag.NewTagRepository(s.provider)
}

func (s *TestSuite) Test_GetAll() {
	type ExpectedReturn struct {
		expectedTags []tag.Tag
		expectedErr  error
	}

	t := s.T()

	cases := map[string]ExpectedReturn{
		"should call provider.GetAll": {
			expectedTags: []tag.Tag{{Id: 1, Name: "Tag 1"}, {Id: 2, Name: "Tag 2"}},
			expectedErr:  nil,
		},
		"should forward error from provider.GetAll": {
			expectedTags: nil,
			expectedErr:  errors.New("error on provider.GetAll"),
		},
	}

	for name, er := range cases {
		s.ResetMocks()
		s.Run(name, func() {
			s.provider.On("GetAll").Return(er.expectedTags, er.expectedErr).Once()

			tags, err := s.repo.GetAll()

			s.provider.AssertNumberOfCalls(t, "GetAll", 1)
			s.Equal(er.expectedTags, tags)
			s.Equal(er.expectedErr, err)
		})
	}
}

func (s *TestSuite) Test_GetById() {
	type ExpectedReturn struct {
		expectedTag tag.Tag
		expectedErr error
	}

	t := s.T()

	cases := map[string]ExpectedReturn{
		"should call provider.GetById": {
			expectedTag: tag.Tag{Id: 1, Name: "Tag 1"},
			expectedErr: nil,
		},
		"should forward error from provider.GetById": {
			expectedTag: tag.Tag{},
			expectedErr: errors.New("error on provider.GetById"),
		},
	}

	for name, er := range cases {
		s.ResetMocks()
		s.Run(name, func() {
			s.provider.On("GetById", 1).Return(er.expectedTag, er.expectedErr).Once()

			tag, err := s.repo.GetById(1)

			s.provider.AssertNumberOfCalls(t, "GetById", 1)
			s.Equal(er.expectedTag, tag)
			s.Equal(er.expectedErr, err)
		})
	}
}

func (s *TestSuite) Test_GetByName() {
	type ExpectedReturn struct {
		expectedTag tag.Tag
		expectedErr error
	}

	t := s.T()

	cases := map[string]ExpectedReturn{
		"should call provider.GetByName": {
			expectedTag: tag.Tag{Id: 1, Name: "Tag 1"},
			expectedErr: nil,
		},
		"should forward error from provider.GetByName": {
			expectedTag: tag.Tag{},
			expectedErr: errors.New("error on provider.GetByName"),
		},
	}

	for name, er := range cases {
		s.ResetMocks()
		s.Run(name, func() {
			s.provider.On("GetByName", "Tag 1").Return(er.expectedTag, er.expectedErr).Once()

			tag, err := s.repo.GetByName("Tag 1")

			s.provider.AssertNumberOfCalls(t, "GetByName", 1)
			s.Equal(er.expectedTag, tag)
			s.Equal(er.expectedErr, err)
		})
	}
}

func (s *TestSuite) Test_Create() {
	type ExpectedReturn struct {
		expectedId  int
		expectedErr error
	}

	t := s.T()

	cases := map[string]ExpectedReturn{
		"should call provider.Create": {
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
			s.provider.On("Create", "Tag 1").Return(er.expectedId, er.expectedErr).Once()

			id, err := s.repo.Create("Tag 1")

			s.provider.AssertNumberOfCalls(t, "Create", 1)
			s.Equal(er.expectedId, id)
			s.Equal(er.expectedErr, err)
		})
	}
}
