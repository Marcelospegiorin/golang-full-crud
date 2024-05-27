package database

import "github.com/Marcelospegiorin/crud-golang-example/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(emaiId string) (*entity.User, error)
}

type ProductInterface interface {
	Create(product *entity.Product) (*entity.Product, error)
	FindAll(page, limit int, sort string) ([]entity.Product, error)
	FindByID(id string) (*entity.Product, error)
	Update(product *entity.Product) error
	Delete(id string) error
}

type CategoryInterface interface {
	Create(category *entity.Category) error
	FindByID(id string) (*entity.Category, error)
	FindAll(sort string) ([]entity.Category, error)
	Delete(id string) error
}
