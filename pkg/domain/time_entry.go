package domain

import (
	"time"
)

type TimeEntryTime struct {
	time.Time
}

func (t *TimeEntryTime) UnmarshalJSON(b []byte) (err error) {
	dateFormat := `"2006-01-02T15:04:05-07:00"`
	dateString := string(b)

	if dateString == "null" {
		t.Time = time.Time{}
		return
	}

	if b[len(b)-2] == 'Z' {
		dateFormat = `"2006-01-02T15:04:05Z"`
	}

	date, err := time.Parse(dateFormat, dateString)
	if err != nil {
		return err
	}

	t.Time = date.In(t.UTC().Location())
	return
}

type TimeEntry struct {
	Id          int           `json:"id" bson:"id,omitempty"`
	Description string        `json:"description" bson:"description,omitempty"`
	StartDate   TimeEntryTime `json:"start" bson:"start,omitempty"`
	EndDate     TimeEntryTime `json:"stop" bson:"stop,omitempty"`
}

type TimeEntryService interface {
	GetCurrent() (*TimeEntry, error)
	List() ([]*TimeEntry, error)
	Create(timeEntry *TimeEntry) (int, error)
	Update(timeEntry *TimeEntry) error
	Delete(id int) error
}

type TimeEntryDB interface {
	GetCurrent() (*TimeEntry, error)
	List() ([]*TimeEntry, error)
	Create(timeEntry *TimeEntry) (int, error)
	Update(timeEntry *TimeEntry) error
	Delete(id int) error
}
