package database

import (
	"github.com/aserafim/pos_golang/09-APIS/internal/entity"
)

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
