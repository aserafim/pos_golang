package main

import (
	"net/http"

	"github.com/aserafim/pos_golang/09_APIs/configs"
	_ "github.com/aserafim/pos_golang/09_APIs/docs"
	"github.com/aserafim/pos_golang/09_APIs/internal/database"
	"github.com/aserafim/pos_golang/09_APIs/internal/entity"
	"github.com/aserafim/pos_golang/09_APIs/internal/webserver/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// @title           User and Products API
// @version         1.0
// @description     API developed as part of the Golang Developer course.
// @termsOfService  http://swagger.io/terms/

// @contact.name   Alefe Serafim
// @contact.url    https://www.linkedin.com/in/alefe-serafim/
// @contact.email  correia.alefe@gmail.com

// @license.name  MIT
// @license.url   https://opensource.org/license/mit

// @host      localhost:8080
// @BasePath  /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
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

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8080/docs/doc.json")))
	http.ListenAndServe(":8080", r)

}
