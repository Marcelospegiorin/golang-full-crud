package database

import (
	"github.com/Marcelospegiorin/crud-golang-example/internal/entity"
	"gorm.io/gorm"
)

type Category struct {
	DB *gorm.DB
}

func NewCategory(db *gorm.DB) *Category {
	return &Category{DB: db}
}

func (c *Category) Create(category *entity.Category) error {
	return c.DB.Create(category).Error
}

func (c *Category) FindByID(id string) (*entity.Category, error) {
	var category entity.Category
	err := c.DB.First(&category, "id = ?", id).Error
	return &category, err
}

func (c *Category) FindAll(sort string) ([]entity.Category, error) {
	var categories []entity.Category
	var err error
	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}
	err = c.DB.Order("created_at " + sort).Find(&categories).Error
	return categories, err
}

func (c *Category) Delete(id string) error {
	category, err := c.FindByID(id)
	if err != nil {
		return err
	}
	return c.DB.Delete(category).Error
}
