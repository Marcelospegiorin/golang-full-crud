package database

import (
	"github.com/Marcelospegiorin/crud-golang-example/internal/entity"
	"gorm.io/gorm"
)

type Product struct {
	DB *gorm.DB
}

func NewProduct(db *gorm.DB) *Product {
	return &Product{DB: db}
}

func (p *Product) Create(product *entity.Product) (*entity.Product, error) {
	result := p.DB.Create(product)
	if result.Error != nil {
		return nil, result.Error
	}
	var productCreated entity.Product
	errDb := p.DB.First(&productCreated, "id = ?", product.ID).Error
	return &productCreated, errDb
}

func (p *Product) FindByID(id string) (*entity.Product, error) {
	var product entity.Product
	err := p.DB.First(&product, "id = ?", id).Error
	return &product, err
}

func (p *Product) Delete(id string) error {
	product, err := p.FindByID(id)
	if err != nil {
		return err
	}
	return p.DB.Delete(product).Error
}

func (p *Product) Update(product *entity.Product) error {
	_, err := p.FindByID(product.ID.String())
	if err != nil {
		return nil
	}
	return p.DB.Save(product).Error
}

func (p *Product) FindAll(page, limit int, sort string) ([]entity.Product, error) {
	var products []entity.Product
	var err error
	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}
	if page != 0 && limit != 0 {
		err = p.DB.Limit(limit).Offset((page - 1) * limit).Order("created_at " + sort).Find(&products).Error
	} else {
		err = p.DB.Order("created_at " + sort).Find(&products).Error
	}
	return products, err
}

func (p *Product) AddCategory(productID string, category *entity.Category) error {
	product, err := p.FindByID(productID)
	if err != nil {
		return err
	}
	product.Categories = append(product.Categories, *category)
	return p.DB.Save(product).Error
}

func (p *Product) RemoveCategory(productID string, categoryID string) error {
	product, err := p.FindByID(productID)
	if err != nil {
		return err
	}
	for i, cat := range product.Categories {
		if cat.ID.String() == categoryID {
			product.Categories = append(product.Categories[:i], product.Categories[i+1:]...)
			break
		}
	}
	return p.DB.Save(product).Error
}
