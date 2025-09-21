package use_case

import (
	"github.com/freit4sdev/go-api/internal/entity"
	"github.com/freit4sdev/go-api/internal/repository"
)

type createCategoryUseCase struct {
	repo repository.ICategoryRepository
}

func NewCreateCategoryUseCase(repo repository.ICategoryRepository) *createCategoryUseCase {
	return &createCategoryUseCase{repo: repo}
}

func (c *createCategoryUseCase) Execute(name string) error {
	category, err := entity.NewCategory(name)
	if err != nil {
		return err
	}

	err = c.repo.Save(category)
	if err != nil {
		return err
	}

	return nil
}
