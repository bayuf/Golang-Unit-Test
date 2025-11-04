package services

import (
	entity "Go-Unit-Test/Go-Mock-Test/Entity"
	repository "Go-Unit-Test/Go-Mock-Test/Repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryServices = CategoryService{Repository: categoryRepository}

func TestCategoryServices_GetNotFound(t *testing.T) {
	// Mock Program
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	category, err := categoryServices.Get("1")
	assert.NotNil(t, err)
	assert.Nil(t, category)
}

func TestCategoryServices_GetFound(t *testing.T) {
	category := entity.Category{
		Id:   "1",
		Name: "Book",
	}

	categoryRepository.Mock.On("FindById", "2").Return(category)

	result, err := categoryServices.Get("2")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, category.Id, result.Id)
	assert.Equal(t, category.Name, result.Name)

}
