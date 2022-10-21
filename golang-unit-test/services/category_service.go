package services

import (
	"errors"
	"golang-unit-test/entity"
	"golang-unit-test/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (categoryService CategoryService) Get(Id string) (*entity.Category, error) {
	category := categoryService.Repository.FindById(Id)
	if category == nil {
		return category, errors.New("Not Found")
	} else {
		return category, nil
	}
}
