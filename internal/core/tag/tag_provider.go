package tag

type TagProvider interface {
	GetAll() ([]Tag, error)
	GetById(id int) (Tag, error)
	GetByName(name string) (Tag, error)
	Create(name string) (id int, err error)
	Delete(id int) error
	UpdateName(id int, name string) error
}
