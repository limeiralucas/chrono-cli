package time_entry

import (
	"time"

	"github.com/limeiralucas/chrono-cli/internal/tag"
)

type TimeEntryProvider interface {
	Create(title string, startDate time.Time) (id int, err error)
	Delete(id int) error
	GetTags(id int) ([]tag.Tag, error)
	UpdateTitle(id int, title string) error
	AddTags(id int, tags []tag.Tag) error
	RemoveTags(id int, tags []tag.Tag) error
}

type TimeEntryRepository struct {
	provider TimeEntryProvider
}

func NewTimeEntryRepository(provider TimeEntryProvider) TimeEntryRepository {
	return TimeEntryRepository{provider: provider}
}

func (t TimeEntryRepository) Create(title string, startDate time.Time) (int, error) {
	return t.provider.Create(title, startDate)
}

func (t TimeEntryRepository) Delete(id int) error {
	return t.provider.Delete(id)
}

func (t TimeEntryRepository) GetTags(id int) ([]tag.Tag, error) {
	return t.provider.GetTags(id)
}

func (t TimeEntryRepository) AddTags(id int, tags []tag.Tag) error {
	return t.provider.AddTags(id, tags)
}

func (t TimeEntryRepository) RemoveTags(id int, tags []tag.Tag) error {
	return t.provider.RemoveTags(id, tags)
}

func (t TimeEntryRepository) UpdateTitle(id int, title string) error {
	return t.provider.UpdateTitle(id, title)
}
