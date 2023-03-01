package toggl

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/limeiralucas/chrono-cli/pkg/domain"
)

const API_HOST = "https://api.track.toggl.com/api/v9"

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var (
	Client HTTPClient = &http.Client{}
)

type TimeEntryProvider struct {
	apiKey string
}

// Create implements domain.TimeEntryDB
func (TimeEntryProvider) Create(timeEntry *domain.TimeEntry) (int, error) {
	panic("unimplemented")
}

// Delete implements domain.TimeEntryDB
func (TimeEntryProvider) Delete(id int) error {
	panic("unimplemented")
}

// Update implements domain.TimeEntryDB
func (TimeEntryProvider) Update(timeEntry *domain.TimeEntry) error {
	panic("unimplemented")
}

func (t *TimeEntryProvider) GetCurrent() (*domain.TimeEntry, error) {
	url := fmt.Sprintf("%s%s", API_HOST, "/me/time_entries/current")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(t.apiKey, "api_token")
	res, err := Client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var timeEntry domain.TimeEntry
	err = json.Unmarshal(body, &timeEntry)
	if err != nil {
		return nil, err
	}

	return &timeEntry, nil
}

func (t *TimeEntryProvider) List() ([]*domain.TimeEntry, error) {
	url := fmt.Sprintf("%s%s", API_HOST, "/me/time_entries")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(t.apiKey, "api_token")
	res, err := Client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var timeEntry []*domain.TimeEntry
	err = json.Unmarshal(body, &timeEntry)
	if err != nil {
		return nil, err
	}

	return timeEntry, nil
}

func NewTimeEntryProvider(apiKey string) TimeEntryProvider {
	return TimeEntryProvider{apiKey: apiKey}
}
