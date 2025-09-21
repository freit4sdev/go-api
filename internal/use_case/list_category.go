package use_case

import (
	"github.com/freit4sdev/go-api/internal/entity"
	"github.com/freit4sdev/go-api/internal/repository"
)

type listCategoryUseCase struct {
	repo repository.ICategoryRepository
}

func NewListCategoryUseCase(repo repository.ICategoryRepository) *listCategoryUseCase {
	return &listCategoryUseCase{repo: repo}
}

func (c *listCategoryUseCase) Execute() ([]*entity.Category, error) {
	categories, err := c.repo.List()
	if err != nil {
		return nil, err
	}

	return categories, nil
}
