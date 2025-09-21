package entity

import (
	"fmt"
	"time"
)

type Category struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewCategory(name string) (*Category, error) {
	category := &Category{
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	
	err := category.IsValid()
	if err != nil {
		return nil, err
	}
	
	return category, nil
}

func (c *Category) IsValid() error {
	if len(c.Name) < 5 {
		return fmt.Errorf("category must have at least 5 characters")
	}
	return nil
}
