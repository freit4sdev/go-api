package use_case

import (
	"log"

	"github.com/freit4sdev/go-api/internal/entity"
)

type createCategoryUseCase struct {
}

func NewCreateCategoryUseCase() *createCategoryUseCase {
	return &createCategoryUseCase{}
}

func (c *createCategoryUseCase) Execute(name string) error {
	category, err := entity.NewCategory(name)
	if err != nil {
		return err
	}
	
	log.Println(category)
	return nil
}
