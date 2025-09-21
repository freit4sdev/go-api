package repository

import "github.com/freit4sdev/go-api/internal/entity"

type categoryRepository struct {
	db []*entity.Category
}

func NewCategoryRepository() *categoryRepository {
	return &categoryRepository{
		db: make([]*entity.Category, 0),
	}
}

func (r *categoryRepository) Save(entity *entity.Category) error {
	r.db = append(r.db, entity)
	return nil
}

func(r *categoryRepository) List() ([]*entity.Category, error) {
	return r.db, nil
}
