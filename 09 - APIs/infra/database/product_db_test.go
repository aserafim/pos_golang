package database

import (
	"fmt"
	"math/rand/v2"
	"testing"

	"github.com/aserafim/pos_golang/09_APIs/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file:memory"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Iphono", 12000.00)
	assert.NoError(t, err)

	productDB := NewProduct(db)
	err = productDB.Create(product)
	assert.NoError(t, err)
}

func TestFindAllProducts(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file:memory2"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})

	for i := 1; i <= 30; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Produto %02d", i), rand.Float32())
		assert.NoError(t, err)
		db.Create(product)
	}

	productDB := NewProduct(db)
	products, err := productDB.FindAll(1, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Produto 01", products[0].Name)
	assert.Equal(t, "Produto 10", products[9].Name)

	products = nil
	products, err = productDB.FindAll(2, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Produto 11", products[0].Name)
	assert.Equal(t, "Produto 20", products[9].Name)

	products = nil
	products, err = productDB.FindAll(3, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Produto 21", products[0].Name)
	assert.Equal(t, "Produto 30", products[9].Name)
}

func TestFindByID(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory3:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}

	db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Produto 01", 5000.00)
	assert.Nil(t, err)
	productDB := NewProduct(db)
	err = productDB.Create(product)

	assert.Nil(t, err)
	productFound, err := productDB.FindByID(product.ID.String())
	assert.Nil(t, err)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)

}

func TestUpdate(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory3:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}

	db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Produto 01", 5000.00)
	assert.Nil(t, err)
	productDB := NewProduct(db)
	err = productDB.Create(product)
	assert.Nil(t, err)
	product.Name = "Produto 02"
	product.Price = 7000.00

	err = productDB.Update(product)
	assert.Nil(t, err)

	productFound, err := productDB.FindByID(product.ID.String())
	assert.Nil(t, err)
	assert.Equal(t, productFound.Name, product.Name)
	assert.Equal(t, productFound.Price, product.Price)
}

func TestDelete(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory3:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}

	db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Produto 01", 5000.00)
	assert.Nil(t, err)
	productDB := NewProduct(db)
	err = productDB.Create(product)
	assert.Nil(t, err)

	err = productDB.Delete(product.ID.String())
	assert.Nil(t, err)

	productFound, err := productDB.FindByID(product.ID.String())
	assert.NotNil(t, err)
	assert.NotEqual(t, product.Name, productFound.Name)
	assert.NotEqual(t, product.Price, productFound.Price)
}

/*
func (p *Product) Delete(id string) error {
	product, err := p.FindByID(id)
	if err != nil {
		return err
	}
	return p.DB.Delete(product).Error
}
*/
