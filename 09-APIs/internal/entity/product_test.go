package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	prod, err := NewProduct("BPhone 7", 7800.00)
	assert.Nil(t, err)
	assert.NotNil(t, prod)
	assert.NotEmpty(t, prod.ID)
	assert.NotEmpty(t, prod.Name)
	assert.Equal(t, "BPhone 7", prod.Name)
	assert.Equal(t, float32(7800.00), prod.Price)
}

/*
func TestProductPrice(t *testing.T) {
	prod, err := NewProduct("BPhone 7", 7800.00)
	assert.Nil(t, err)
	assert.NotEmpty(t, prod.Price)
	assert.Greater(t, prod.Price, float32(0.0))

}
*/

func TestProductWhenNameIsRequired(t *testing.T) {
	p, err := NewProduct("", 50.0)
	assert.Nil(t, p)
	assert.Equal(t, ErrNameIsRequired, err)
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	p, err := NewProduct("Teste", 0)
	assert.Nil(t, p)
	assert.Equal(t, ErrPriceIsRequired, err)
}

func TestProductWhenPriceIsInvalid(t *testing.T) {
	p, err := NewProduct("Teste", -10)
	assert.Nil(t, p)
	assert.Equal(t, ErrInvalidPrice, err)
}
	
func TestProductValidate(t *testing.T) {
	p, err := NewProduct("Prod 1", 55.00)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.Nil(t, p.Validate())
}
