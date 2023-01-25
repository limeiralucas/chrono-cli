package tag

type TagRepository interface {
	Create(name string) (ref any, err error)
	Delete(ref any) error
	UpdateName(name string) error
}
