package repository

type ICategoryRepository interface {
	Save(name string) error
	
}
