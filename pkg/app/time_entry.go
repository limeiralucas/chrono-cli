package app

import (
	"time"

	"github.com/limeiralucas/chrono-cli/pkg/domain"
)

type timeEntryService struct {
	DB domain.TimeEntryDB
}

// Create implements domain.TimeEntryService
func (ts timeEntryService) Create(timeEntry *domain.TimeEntry) (int, error) {
	return ts.DB.Create(timeEntry)
}

// Update implements domain.TimeEntryService
func (ts timeEntryService) Update(timeEntry *domain.TimeEntry) error {
	return ts.DB.Update(timeEntry)
}

// Delete implements domain.TimeEntryService
func (ts timeEntryService) Delete(id int) error {
	return ts.DB.Delete(id)
}

// Get implements domain.TimeEntryService
func (ts timeEntryService) GetCurrent() (*domain.TimeEntry, error) {
	return ts.DB.GetCurrent()
}

// List implements domain.TimeEntryService
func (ts timeEntryService) List(startTime time.Time, endTime time.Time) ([]*domain.TimeEntry, error) {
	return ts.DB.List(startTime, endTime)
}

func (ts timeEntryService) ElapsedTimeByDay(startTime time.Time, endTime time.Time) (map[string]float32, error) {
	elapsedTime := map[string]float32{}
	entries, err := ts.List(startTime, endTime)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		day := entry.StartDate.Format("02/01")
		elapsed, ok := elapsedTime[day]
		if !ok {
			elapsed = 0
		}
		elapsedTime[day] = elapsed + float32(entry.Duration)/60/60
	}

	return elapsedTime, nil
}

func (ts timeEntryService) TimeReport(startTime time.Time, endTime time.Time) (map[string]domain.TagReport, error) {
	report := map[string]domain.TagReport{}

	entries, err := ts.List(startTime, endTime)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		day := entry.StartDate.Format("02/01")
		dayReport, ok := report[day]
		if !ok {
			dayReport = domain.TagReport{}
		}
		label := "[other]"
		tags := entry.Tags
		if len(tags) > 0 {
			label = tags[0]
		}

		tagReport, ok := dayReport[label]
		if !ok {
			tagReport = []*domain.TimeEntry{}
		}
		tagReport = append(tagReport, entry)
		dayReport[label] = tagReport

		report[day] = dayReport
	}

	return report, nil
}

func NewTimeEntryService(db domain.TimeEntryDB) domain.TimeEntryService {
	return timeEntryService{
		DB: db,
	}
}
