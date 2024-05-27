package entity

import (
	"time"

	"github.com/Marcelospegiorin/crud-golang-example/pkg/entity"
)

type Category struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type ProductCategory struct {
	ProductID  entity.ID `gorm:"primaryKey" json:"product_id"`
	CategoryID entity.ID `gorm:"primaryKey" json:"category_id"`
}

func NewCategory(name string) (*Category, error) {
	category := &Category{
		ID:        entity.NewID(),
		Name:      name,
		CreatedAt: time.Now(),
	}
	return category, nil
}
