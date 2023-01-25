package time_entry

import (
	"time"

	"github.com/limeiralucas/chrono-cli/internal/tag"
)

type TimeEntryProvider interface {
	Create(title string, startDate time.Time, tags []tag.Tag) (id int, err error)
	Delete(id int) error
	Update(title string, tags []tag.Tag)
	GetTags(id int) ([]tag.Tag, error)
}

type TimeEntryRepository struct {
	provider TimeEntryProvider
}

func NewTimeEntryRepository(provider TimeEntryProvider) TimeEntryRepository {
	return TimeEntryRepository{provider: provider}
}

func (t TimeEntryRepository) Create(title string, startDate time.Time) (int, error) {
	return t.provider.Create(title, startDate, nil)
}

func (t TimeEntryRepository) Delete(id int) error {
	return t.provider.Delete(id)
}

func (t TimeEntryRepository) GetTags(id int) ([]tag.Tag, error) {
	return t.provider.GetTags(id)
}
