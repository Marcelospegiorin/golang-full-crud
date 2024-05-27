package entity

import (
	"errors"
	"time"

	"github.com/Marcelospegiorin/crud-golang-example/pkg/entity"
)

type Product struct {
	ID         entity.ID  `json:"id"`
	Name       string     `json:"name"`
	Price      float64    `json:"price"`
	CreatedAt  time.Time  `json:"created_at"`
	Categories []Category `gorm:"many2many:product_categories;" json:"categories,omitempty"`
}

var (
	ErrIDIsRequired    = errors.New("id is required")
	ErrIDInvalid       = errors.New("invalid id")
	ErrPriceIsRequired = errors.New("price is required")
	ErrInvalidPrice    = errors.New("invalid price")
	ErrNameIsRequired  = errors.New("name required")
)

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return ErrIDIsRequired
	}
	if p.Name == "" {
		return ErrNameIsRequired
	}
	if p.Price == 0 {
		return ErrPriceIsRequired
	}
	if p.Price < 0 {
		return ErrInvalidPrice
	}
	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return ErrIDInvalid
	}
	return nil
}

func NewProduct(name string, price float64) (*Product, error) {
	product := &Product{
		ID:        entity.NewID(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}
	err := product.Validate()
	if err != nil {
		return nil, err
	}
	return product, nil
}
