package tag

import "github.com/limeiralucas/chrono-cli/internal/core/tag"

type TagRepository struct {
	provider tag.TagProvider
}

func NewTagRepository(provider tag.TagProvider) TagRepository {
	return TagRepository{provider: provider}
}

func (t *TagRepository) GetAll() ([]tag.Tag, error) {
	return t.provider.GetAll()
}

func (t *TagRepository) GetById(id int) (tag.Tag, error) {
	return t.provider.GetById(id)
}

func (t *TagRepository) GetByName(name string) (tag.Tag, error) {
	return t.provider.GetByName(name)
}

func (t *TagRepository) Create(name string) (int, error) {
	return t.provider.Create(name)
}
