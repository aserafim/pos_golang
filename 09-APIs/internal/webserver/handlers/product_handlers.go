package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/aserafim/pos_golang/09-APIS/internal/dto"
	"github.com/aserafim/pos_golang/09-APIS/internal/entity"
	"github.com/aserafim/pos_golang/09-APIS/internal/infra/database"
)

type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product dto.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	p, err := entity.NewProduct(product.Name, float32(product.Price))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	err = h.ProductDB.Create(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
