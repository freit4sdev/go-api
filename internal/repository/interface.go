package repository

import "github.com/freit4sdev/go-api/internal/entity"

type ICategoryRepository interface {
	Save(category *entity.Category) error
	List() ([]*entity.Category, error)
}
