package app

import (
	"fmt"
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
	tag := ""
	elapsedTime := map[string]float32{}
	entries, err := ts.List(startTime, endTime)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if len(entry.Tags) > 0 {
			tag = fmt.Sprintf(" (%s)", entry.Tags[0])
		} else {
			tag = ""
		}

		day := entry.StartDate.Format("02/01")
		key := fmt.Sprintf("%s%s", day, tag)
		elapsed, ok := elapsedTime[day]
		if !ok {
			elapsed = 0
		}
		elapsedTime[key] = elapsed + float32(entry.Duration)/60/60
	}

	return elapsedTime, nil
}

func NewTimeEntryService(db domain.TimeEntryDB) domain.TimeEntryService {
	return timeEntryService{
		DB: db,
	}
}
