package tag

type TagProvider interface {
	GetAll() ([]Tag, error)
	GetById(id int) (Tag, error)
	GetByName(name string) (Tag, error)
	Create(name string) (id int, err error)
	Delete(id int) error
	UpdateName(id int, name string) error
}

type TagRepository struct {
	provider TagProvider
}

func NewTagRepository(provider TagProvider) TagRepository {
	return TagRepository{provider: provider}
}

func (t TagRepository) GetAll() ([]Tag, error) {
	return t.provider.GetAll()
}

func (t TagRepository) GetById(id int) (Tag, error) {
	return t.provider.GetById(id)
}

func (t TagRepository) GetByName(name string) (Tag, error) {
	return t.provider.GetByName(name)
}

func (t TagRepository) Create(name string) (int, error) {
	return t.provider.Create(name)
}
