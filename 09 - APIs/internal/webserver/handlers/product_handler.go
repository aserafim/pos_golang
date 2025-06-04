package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/aserafim/pos_golang/09_APIs/internal/database"
	"github.com/aserafim/pos_golang/09_APIs/internal/dto"
	"github.com/aserafim/pos_golang/09_APIs/internal/entity"
	"github.com/go-chi/chi"
)

// Cria uma struct de Handler
type ProductHandler struct {
	ProductDB database.ProductInterface
}

// Cria um construtor para ProductHandler
func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}

// Cria uma Handle Func para criar um novo Produto
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var productInput dto.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&productInput)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	product, err := entity.NewProduct(productInput.Name, productInput.Price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.ProductDB.Create(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// Cria um Handler para GetById
func (h *ProductHandler) GetProductById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	product, err := h.ProductDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}
