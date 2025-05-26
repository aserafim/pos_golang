package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	product, err := NewProduct("Iphono", 12000)
	assert.Nil(t, err)
	assert.NotEmpty(t, product)
	assert.Equal(t, "Iphono", product.Name)
	assert.Equal(t, float32(12000), product.Price)
	assert.NotEmpty(t, product.ID)
	assert.True(t, product.Price > 0)
	assert.False(t, product.Price <= 0)
}

func TestWhenPriceIsInvalid(t *testing.T) {
	prod, err := NewProduct("Samsungo", -8999)
	assert.Nil(t, prod)
	assert.Equal(t, ErrInvalidPrice, err)
}

func TestWhenNameIsEmpty(t *testing.T) {
	prod, err := NewProduct("", 7000)
	assert.Nil(t, prod)
	assert.Equal(t, ErrNameIsRequired, err)
}
