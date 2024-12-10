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
//	assert.Equal(t, product.ID.String(), productFound.ID.String())
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.CreatedAt.Unix(), productFound.CreatedAt.Unix())
	assert.Equal(t, product.Price, productFound.Price)
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

	// ID        entity.ID `json:"id"`
	// Name      string    `json:"name"`
	// Price     float32   `json:"price"`
	// CreatedAt time.Time `json:"created_at"`

	productFound, err := productDB.FindByID(product.ID.String())
	assert.Nil(t, err)
	//assert.Equal(t, product.ID.String(), productFound.ID.String())
	//t.Log("product: ", product.ID.String())
	//t.Log("productFound: ", productFound.ID.String())
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
	assert.NotNil(t, product.CreatedAt.Unix(), productFound.CreatedAt.Unix())

	//product_db_test.go:70: product: 		f4b9ab5e-5039-41d1-81ad-c3ce52507972
	//product_db_test.go:71: productFound:  f4b9ab5e-5039-41d1-81ad-c3ce52507972
}

