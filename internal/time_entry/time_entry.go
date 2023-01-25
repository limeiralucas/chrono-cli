package time_entry

import (
	"time"

	"github.com/limeiralucas/chrono-cli/internal/tag"
)

type TimeEntry struct {
	title     string
	startedAt time.Time
	endedAt   time.Time
	tags      []tag.Tag
}
