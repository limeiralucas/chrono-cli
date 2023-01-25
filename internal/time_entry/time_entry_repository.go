package time_entry

import "time"

type TimeEntryRepository interface {
	Create(title string, startDate time.Time) (ref any, err error)
	Delete(ref any) error
	UpdateTitle(title string) error
	AddTag(title string) error
	RemoveTag(title string) error
}
