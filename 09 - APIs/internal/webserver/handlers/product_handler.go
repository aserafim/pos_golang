package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/aserafim/pos_golang/09_APIs/internal/database"
	"github.com/aserafim/pos_golang/09_APIs/internal/dto"
	"github.com/aserafim/pos_golang/09_APIs/internal/entity"
	entityPkg "github.com/aserafim/pos_golang/09_APIs/pkg/entity"
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

// Create Product 		godoc
// @Summary		Create product
// @Description	Create a new product
// @Tags		products
// @Accept		json
// @Produce		json
// @Param		request		body		dto.CreateProductInput		true		"product request"
// @Success 	201
// @Failure		400		{object}	Error
// @Failure		500		{object}	Error
// @Router		/products	[post]
// @Security ApiKeyAuth
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var productInput dto.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&productInput)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	product, err := entity.NewProduct(productInput.Name, productInput.Price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	err = h.ProductDB.Create(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// Get Product 		godoc
// @Summary		Get product
// @Description	Get product by id
// @Tags		products
// @Accept		json
// @Produce		json
// @Param		id		path	string	true	"Product id" Format(uuid)
// @Success 	200		{object}		entity.Product
// @Failure		404		{object}	Error
// @Failure		500		{object}	Error
// @Router		/products/{id}	[get]
// @Security ApiKeyAuth
func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: "id cannot be empty"}
		json.NewEncoder(w).Encode(error)
		return
	}
	product, err := h.ProductDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

// Get Product 		godoc
// @Summary		List products
// @Description	List all products
// @Tags		products
// @Accept		json
// @Produce		json
// @Param		page		query	string	false	"page number"
// @Param		limit		query	string	false	"limit"
// @Success 	200		{array}		entity.Product
// @Failure		404		{object}	Error
// @Failure		500		{object}	Error
// @Router		/products	[get]
// @Security ApiKeyAuth
func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 0
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 0
	}

	sort := r.URL.Query().Get("sort")

	products, err := h.ProductDB.FindAll(pageInt, limitInt, sort)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)

}

// Update Product 		godoc
// @Summary		Update product
// @Description	Update product information
// @Tags		products
// @Accept		json
// @Produce		json
// @Param		id			path	string						true	"Product id" Format(uuid)
// @Param		request		body	dto.CreateProductInput		true		"product request"
// @Success 	200
// @Failure		400		{object}	Error
// @Failure		401		{object}	Error
// @Failure		500		{object}	Error
// @Router		/products/{id}	[put]
// @Security ApiKeyAuth
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: "id cannot be empty"}
		json.NewEncoder(w).Encode(error)
		return
	}
	var product entity.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	product.ID, err = entityPkg.ParseID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	_, err = h.ProductDB.FindByID(id)
	if err != nil {
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		w.WriteHeader(http.StatusNotFound)
	}

	err = h.ProductDB.Update(&product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	w.WriteHeader(http.StatusOK)

}

// Delete Product 		godoc
// @Summary		Delete product
// @Description	Delete product by id
// @Tags		products
// @Accept		json
// @Produce		json
// @Param		id		path	string	true	"Product id" Format(uuid)
// @Success 	200
// @Failure		400		{object}	Error
// @Failure		404		{object}	Error
// @Failure		500		{object}	Error
// @Router		/products/{id}	[delete]
// @Security ApiKeyAuth
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: "id cannot be empty"}
		json.NewEncoder(w).Encode(error)
		return
	}
	_, err := h.ProductDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	err = h.ProductDB.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
	}
	w.WriteHeader(http.StatusOK)
}
