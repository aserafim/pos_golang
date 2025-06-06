package database

import (
	"github.com/aserafim/pos_golang/09_APIs/internal/entity"
	"gorm.io/gorm"
)

// Um user atrelado
// a um banco de dados
type User struct {
	DB *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{DB: db}
}

// Implementa funcao que cria o user
// no banco e retorna o erro se houver
func (u *User) Create(user *entity.User) error {
	return u.DB.Create(user).Error
}

func (u *User) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	if err := u.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
