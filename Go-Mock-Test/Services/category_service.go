package services

import (
	entity "Go-Unit-Test/Go-Mock-Test/Entity"
	repository "Go-Unit-Test/Go-Mock-Test/Repository"
	"errors"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (services CategoryService) Get(id string) (*entity.Category, error) {
	category := services.Repository.FindById(id)
	if category == nil {
		return category, errors.New("category tidak ditemukan")
	} else {
		return category, nil
	}
}
