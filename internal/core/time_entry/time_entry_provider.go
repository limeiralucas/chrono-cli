package time_entry

import (
	"time"

	"github.com/limeiralucas/chrono-cli/internal/core/tag"
)

type TimeEntryProvider interface {
	Create(title string, startDate time.Time) (id int, err error)
	Delete(id int) error
	GetTags(id int) ([]tag.Tag, error)
	UpdateTitle(id int, title string) error
	AddTags(id int, tags []tag.Tag) error
	RemoveTags(id int, tags []tag.Tag) error
}
