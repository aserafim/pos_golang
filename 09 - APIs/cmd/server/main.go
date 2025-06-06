package main

import (
	"net/http"

	"github.com/aserafim/pos_golang/09_APIs/configs"
	"github.com/aserafim/pos_golang/09_APIs/internal/database"
	"github.com/aserafim/pos_golang/09_APIs/internal/entity"
	"github.com/aserafim/pos_golang/09_APIs/internal/webserver/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	//Carrega as configurações
	configs, err := configs.LoadConfig(".")
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
	userDB := database.NewUser(db)

	// Cria um productHandler
	productHandler := handlers.NewProductHandler(productDB)
	userHandler := handlers.NewUserHandler(userDB, configs.TokenAuth, configs.JwtExpiresIn)

	// Cria uma rota com o Chi
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/products", productHandler.CreateProduct)
	r.Get("/products/{id}", productHandler.GetProduct)
	r.Put("/products/{id}", productHandler.UpdateProduct)
	r.Delete("/products/{id}", productHandler.DeleteProduct)
	r.Get("/products", productHandler.GetProducts)
	r.Post("/users", userHandler.CreateUser)
	r.Post("/users/get_token", userHandler.GetJwt)
	http.ListenAndServe(":8080", r)

}
