package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	p, err := NewProduct("Produto 1", 10.0)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.Equal(t, "Produto 1", p.Name)
	assert.Equal(t, 10.0, p.Price)
}

func TestNewProductWhenNameIsRequired(t *testing.T) {
	p, err := NewProduct("", 10.0)
	assert.Nil(t, p)
	assert.Equal(t, ErrNameIsRequired, err)
}

func TestNewProductWhenPriceIsRequired(t *testing.T) {
	p, err := NewProduct("Produto 1", 0)
	assert.Nil(t, p)
	assert.NotNil(t, err)
	assert.Equal(t, ErrPriceIsRequired, err)
}

func TestNewProductWhenPriceIsInvalid(t *testing.T) {
	p, err := NewProduct("Produto 1", -1)
	assert.Nil(t, p)
	assert.NotNil(t, err)
	assert.Equal(t, ErrInvalidPrice, err)
}
