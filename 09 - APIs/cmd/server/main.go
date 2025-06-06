package main

import (
	"net/http"

	"github.com/aserafim/pos_golang/09_APIs/configs"
	"github.com/aserafim/pos_golang/09_APIs/internal/database"
	"github.com/aserafim/pos_golang/09_APIs/internal/entity"
	"github.com/aserafim/pos_golang/09_APIs/internal/webserver/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
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

	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(configs.TokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Post("/", productHandler.CreateProduct)
		r.Get("/", productHandler.GetProducts)
		r.Get("/{id}", productHandler.GetProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	r.Route("/users", func(r chi.Router) {
		r.Post("/", userHandler.CreateUser)
		r.Post("/get_token", userHandler.GetJwt)
	})

	http.ListenAndServe(":8080", r)

}
