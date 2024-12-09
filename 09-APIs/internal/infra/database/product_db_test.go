package database

import (
	"testing"

	"github.com/aserafim/pos_golang/09-APIS/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//ID        entity.ID `json:"id"`
//Name      string    `json:"name"`
//Price     float32   `json:"price"`
//CreatedAt time.Time `json:"created_at"`

func TestCreateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})
	product, _ := entity.NewProduct("IPhono 12", 12999.00)
	productDB := NewProduct(db)

	err = productDB.Create(product)
	assert.Nil(t, err)

	var productFound entity.Product
	err = db.First(&productFound, "id = ?", product.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.CreatedAt.Unix(), productFound.CreatedAt.Unix())
	assert.Equal(t, product.Price, productFound.Price)
}

// func (p *Product) FindByID(id string) (*entity.Product, error) {
// 	var product entity.Product
// 	err := p.DB.First(&product, "id = ?", id).Error
// 	return &product, err
// }
