package entity

import (
	"errors"
	"time"

	"github.com/aserafim/pos_golang/09_APIs/pkg/entity"
)

var (
	ErrIDIsRequired   = errors.New("id is required")
	ErrNameIsRequired = errors.New("name is required")
	ErrInvalidPrice   = errors.New("invalida price")
)

type Product struct {
	ID        entity.ID `json:"ID"`
	Name      string    `json:"name"`
	Price     float32   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func NewProduct(name string, price float32) (*Product, error) {
	product := &Product{
		ID:        entity.NewID(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}

	err := product.Validate()
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return ErrIDIsRequired
	}

	if p.Name == "" {
		return ErrNameIsRequired
	}

	if p.Price <= 0 {
		return ErrInvalidPrice
	}
	return nil
}
