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

func NewTimeEntryService(db domain.TimeEntryDB) domain.TimeEntryService {
	return timeEntryService{
		DB: db,
	}
}
