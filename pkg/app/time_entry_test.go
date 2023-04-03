package app

import (
	"testing"
	"time"

	"github.com/limeiralucas/chrono-cli/mocks"
	"github.com/limeiralucas/chrono-cli/pkg/domain"
	"github.com/stretchr/testify/suite"
)

type TimeEntryServiceSuite struct {
	suite.Suite
	db      domain.TimeEntryDB
	service domain.TimeEntryService
}

func Test_TimeEntryService(t *testing.T) {
	suite.Run(t, new(TimeEntryServiceSuite))
}

func (s *TimeEntryServiceSuite) SetupSuite() {
	s.db = mocks.NewTimeEntryDB(s.T())
	s.service = NewTimeEntryService(s.db)
}

func (s *TimeEntryServiceSuite) Test_TimeReport() {
	startTime := time.Date(2023, 4, 3, 0, 0, 0, 0, time.UTC)
	endTime := time.Date(2023, 4, 8, 0, 0, 0, 0, time.UTC).Add(-time.Nanosecond)

	timeEntries := []*domain.TimeEntry{
		{
			Id:          0,
			Description: "Time Entry 0",
			StartDate:   domain.TimeEntryTime{Time: startTime},
			EndDate:     domain.TimeEntryTime{Time: startTime.Add(time.Hour)},
			Duration:    3600,
			Tags:        []string{"reading"},
		},
		{
			Id:          1,
			Description: "Time Entry 1",
			StartDate:   domain.TimeEntryTime{Time: startTime},
			EndDate:     domain.TimeEntryTime{Time: startTime.Add(time.Hour * 2)},
			Duration:    3600 * 2,
			Tags:        []string{"tech"},
		},
		{
			Id:          2,
			Description: "Time Entry 2",
			StartDate:   domain.TimeEntryTime{Time: startTime},
			EndDate:     domain.TimeEntryTime{Time: startTime.Add(time.Hour * 3)},
			Duration:    3600 * 3,
			Tags:        []string{"tech"},
		},
		{
			Id:          3,
			Description: "Time Entry 3",
			StartDate:   domain.TimeEntryTime{Time: startTime},
			EndDate:     domain.TimeEntryTime{Time: startTime.Add(time.Hour)},
			Duration:    3600,
			Tags:        nil,
		},
		{
			Id:          4,
			Description: "Time Entry 4",
			StartDate:   domain.TimeEntryTime{Time: endTime.Add(time.Hour * -3)},
			EndDate:     domain.TimeEntryTime{Time: endTime},
			Duration:    3600 * 3,
			Tags:        []string{"break"},
		},
	}

	s.db.(*mocks.TimeEntryDB).Mock.On("List", startTime, endTime).Return(timeEntries, nil)
	report, err := s.service.TimeReport(startTime, endTime)
	s.Nil(err)

	// Only 2 days reported
	s.Len(report, 2)

	// Contains day 03/04 with 3 tags
	dayReport, ok := report["03/04"]
	s.True(ok)
	s.Len(dayReport, 3)

	tagReport, ok := dayReport["[other]"]
	s.True(ok)
	s.Equal(tagReport, timeEntries[3:4])

	tagReport, ok = dayReport["tech"]
	s.True(ok)
	s.Equal(tagReport, timeEntries[1:3])

	tagReport, ok = dayReport["reading"]
	s.True(ok)
	s.Equal(tagReport, timeEntries[0:1])

	// Contains day 07/04 with 1 tags
	dayReport, ok = report["07/04"]
	s.True(ok)
	s.Len(dayReport, 1)

	tagReport, ok = dayReport["break"]
	s.True(ok)
	s.Equal(tagReport, timeEntries[4:])
}
