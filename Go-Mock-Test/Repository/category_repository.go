package repository

import entity "Go-Unit-Test/Go-Mock-Test/Entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
