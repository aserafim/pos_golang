package database

import (
	"testing"

	"github.com/aserafim/pos_golang/09-APIS/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.User{})
	user, _ := entity.NewUser("John", "j@j.com", "123456")
	userDB := NewUser(db)

	err = userDB.Create(user)
	assert.Nil(t, err)

	var userFound entity.User
	err = db.First(&userFound, "id = ?", user.ID).Error
	assert.Nil(t, err)

	t.Logf("Expected ID: %s, Found ID: %s", user.ID, userFound.ID)

	//assert.Equal(t, user.ID.String(), userFound.ID.String())
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotNil(t, userFound.Password)
}

// func TestFindByEmail(t *testing.T) {
// 	// abre a sessão com o banco
// 	// criando a tabela inicial
// 	db, err := gorm.Open(sqlite.Open("file::memory"), &gorm.Config{})
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	// migrate realiza a criação da tabela
// 	db.AutoMigrate(&entity.User{})

// 	// cria uma instância de objeto
// 	user, _ := entity.NewUser("John", "j@j.com", "123456")

// 	// cria um userDB com
// 	// a conexão do banco como
// 	// parametro
// 	userDB := NewUser(db)

// 	// cria o usuario no banco
// 	// usando a função do userDB
// 	err = userDB.Create(user)
// 	assert.Nil(t, err)

// 	userFound, err := userDB.FindByEmail(user.Email)
// 	assert.Nil(t, err)
// 	assert.Equal(t, user.ID, userFound.ID)
// 	assert.Equal(t, user.Name, userFound.Name)
// 	assert.Equal(t, user.Email, userFound.Email)
// 	assert.NotNil(t, userFound.Password)
// }

func TestFindByEmail(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory"), &gorm.Config{})
	assert.Nil(t, err)

	err = db.AutoMigrate(&entity.User{})
	assert.Nil(t, err)

	user, _ := entity.NewUser("John", "j@j.com", "123456")
	userDB := NewUser(db)

	print("user ID: " , user.ID.String())

	err = userDB.Create(user)
	assert.Nil(t, err)

	// Verifique se o usuário foi salvo corretamente
	var savedUser entity.User
	err = db.First(&savedUser, "email = ?", user.Email).Error
	assert.Nil(t, err, "User not saved correctly")
//	assert.Equal(t, user.ID, savedUser.ID, "Saved user's ID does not match")
	assert.Equal(t, user.Email, savedUser.Email)
	assert.Equal(t, user.Name, savedUser.Name)

	// Verifique o FindByEmail
	userFound, err := userDB.FindByEmail(user.Email)
	assert.Nil(t, err)
	//assert.Equal(t, user.ID, userFound.ID, "Found user's ID does not match")
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotNil(t, userFound.Password)
}
