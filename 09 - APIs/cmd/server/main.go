package main

import (
	"net/http"

	"github.com/aserafim/pos_golang/09_APIs/configs"
	"github.com/aserafim/pos_golang/09_APIs/internal/database"
	"github.com/aserafim/pos_golang/09_APIs/internal/entity"
	"github.com/aserafim/pos_golang/09_APIs/internal/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	//Carrega as configurações
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	// Cria o banco de dados
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Faz o "migrate" das tabelas
	db.AutoMigrate(&entity.Product{}, &entity.User{})

	// Cria um product DB
	productDB := database.NewProduct(db)

	// Cria um productHandler
	productHandler := handlers.NewProductHandler(productDB)

	// Executa o servidor
	http.HandleFunc("/products", productHandler.CreateProduct)
	http.ListenAndServe(":8080", nil)

}
