package time_entry

import (
	"errors"
	"testing"
	"time"

	"github.com/limeiralucas/chrono-cli/internal/tag"
	"github.com/limeiralucas/chrono-cli/mocks"
	"github.com/stretchr/testify/assert"
)

func TestTimeEntryRepository_Create(t *testing.T) {
	provider := mocks.TimeEntryProvider{}
	repo := NewTimeEntryRepository(&provider)
	now := time.Now()

	testCases := map[string]func(*testing.T){
		"should call provider.Create and return id": func(t *testing.T) {
			provider.On("Create", "entry", now, []tag.Tag(nil)).Return(1, nil).Once()

			id, err := repo.Create("entry", now)

			provider.AssertNumberOfCalls(t, "Create", 1)
			assert.Equal(t, 1, id)
			assert.Nil(t, err)
		},
		"should forward error from provider.Create": func(t *testing.T) {
			provider.On("Create", "entry", now, []tag.Tag(nil)).Return(0, errors.New("error on provider")).Once()

			id, err := repo.Create("entry", now)

			assert.Equal(t, 0, id)
			assert.Error(t, err, "error on provider")
		},
		"should call provider.Delete": func(t *testing.T) {
			provider.On("Delete", 1).Return(nil).Once()

			err := repo.Delete(1)

			provider.AssertNumberOfCalls(t, "Delete", 1)
			assert.Nil(t, err)
		},
		"should forward error from provider.Delete": func(t *testing.T) {
			provider.On("Delete", 1).Return(errors.New("error on provider")).Once()

			err := repo.Delete(1)

			assert.Error(t, err, "error on provider")
		},
		"should call provider.GetTags and return tags": func(t *testing.T) {
			expectedTags := []tag.Tag{
				{Id: 1, Name: "Tag 1"},
				{Id: 2, Name: "Tag 2"},
			}
			provider.On("GetTags", 1).Return(expectedTags, nil).Once()

			tags, err := repo.GetTags(1)

			provider.AssertNumberOfCalls(t, "GetTags", 1)
			assert.Equal(t, expectedTags, tags)
			assert.Nil(t, err)
		},
		"should forward error from provider.GetTags": func(t *testing.T) {
			provider.On("GetTags", 1).Return(nil, errors.New("error on provider")).Once()

			_, err := repo.GetTags(1)

			assert.Error(t, err, "error on provider")
		},
		"should call provider.AddTags": func(t *testing.T) {
			newTags := []tag.Tag{
				{Id: 1, Name: "Tag 1"},
				{Id: 2, Name: "Tag 2"},
			}
			provider.On("AddTags", 1, newTags).Return(nil).Once()

			err := repo.AddTags(1, newTags)

			provider.AssertNumberOfCalls(t, "AddTags", 1)
			assert.Nil(t, err)
		},
		"should forward error from provider.AddTags": func(t *testing.T) {
			provider.On("AddTags", 1, []tag.Tag(nil)).Return(errors.New("error on provider")).Once()

			err := repo.AddTags(1, nil)

			assert.Error(t, err, "error on provider")
		},
		"should call provider.RemoveTags": func(t *testing.T) {
			tagsToBeRemove := []tag.Tag{
				{Id: 1, Name: "Tag 1"},
			}
			provider.On("RemoveTags", 1, tagsToBeRemove).Return(nil).Once()

			err := repo.RemoveTags(1, tagsToBeRemove)

			provider.AssertNumberOfCalls(t, "RemoveTags", 1)
			assert.Nil(t, err)
		},

		"should forward error from provider.RemoveTags": func(t *testing.T) {
			provider.On("RemoveTags", 1, []tag.Tag(nil)).Return(errors.New("error on provider")).Once()

			err := repo.RemoveTags(1, nil)

			assert.Error(t, err, "error on provider")
		},
		"should call provider.UpdateTitle": func(t *testing.T) {
			provider.On("UpdateTitle", 1, "New title").Return(nil).Once()

			err := repo.UpdateTitle(1, "New title")

			assert.Nil(t, err)
			provider.AssertNumberOfCalls(t, "UpdateTitle", 1)
		},
		"should forward error from provider.UpdateTitle": func(t *testing.T) {
			provider.On("UpdateTitle", 1, "").Return(errors.New("error on provider")).Once()

			err := repo.UpdateTitle(1, "")

			assert.Error(t, err, "error on provider")
		},
	}

	for name, run := range testCases {
		t.Run(name, func(t *testing.T) {
			run(t)
		})
	}
}