package main

import (
	"net/http"

	"github.com/aserafim/pos_golang/09-APIS/configs"
	"github.com/aserafim/pos_golang/09-APIS/internal/entity"
	"github.com/aserafim/pos_golang/09-APIS/internal/infra/database"
	"github.com/aserafim/pos_golang/09-APIS/internal/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	_, err := configs.LoadConfig(".")

	if err != nil {
		println(err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})
	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	http.HandleFunc("/products", productHandler.CreateProduct)
	http.ListenAndServe(":8080", nil)
}
