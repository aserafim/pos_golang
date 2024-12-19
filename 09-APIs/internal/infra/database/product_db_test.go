package database

import (
	"fmt"
	"math/rand/v2"
	"testing"

	"github.com/aserafim/pos_golang/09-APIS/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

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
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.CreatedAt.Unix(), productFound.CreatedAt.Unix())
	assert.Equal(t, product.Price, productFound.Price)
}

func TestFindAllProducts(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})
	for i := 1; i < 24; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float32()*100)
		assert.NoError(t, err)
		db.Create(product)
	}

	productDB := NewProduct(db)
	products, err := productDB.FindAll(1, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 10", products[9].Name)

	products, err = productDB.FindAll(2, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 11", products[0].Name)
	assert.Equal(t, "Product 20", products[9].Name)

}

func TestFindByID(t *testing.T) {
	// abre a sessão com o banco
	// criando a tabela inicial
	db, err := gorm.Open(sqlite.Open("file::memory"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	// migrate realiza a criação da tabela
	db.AutoMigrate(&entity.Product{})

	// cria uma instância de objeto
	product, _ := entity.NewProduct("IPhono 12", 12999.00)

	// cria um userDB com
	// a conexão do banco como
	// parametro
	productDB := NewProduct(db)

	// cria o usuario no banco
	// usando a função do userDB
	err = productDB.Create(product)
	assert.Nil(t, err)

	productFound, err := productDB.FindByID(product.ID.String())
	assert.Nil(t, err)

	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
	assert.NotNil(t, product.CreatedAt.Unix(), productFound.CreatedAt.Unix())

}

func TestUpdateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Product 1", 10.00)
	assert.NoError(t, err)
	db.Create(product)

	productDB := NewProduct(db)
	product.Name = "Product 2"
	err = productDB.Update(product)
	assert.NoError(t, err)
	product, err = productDB.FindByID(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, "Product 2", product.Name)
}

func TestDeleteProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Delete 1", 10.00)
	assert.NoError(t, err)
	db.Create(product)
	productDB := NewProduct(db)

	err = productDB.Delete(product.ID.String())
	assert.NoError(t, err)

	_, err = productDB.FindByID(product.ID.String())
	assert.Error(t, err)
	//assert.Nil(t, product)
	assert.Equal(t, gorm.ErrRecordNotFound, err)


}
