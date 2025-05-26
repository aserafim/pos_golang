package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Alefe", "a@a.com", "1234")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "Alefe", user.Name)
	assert.Equal(t, "a@a.com", user.Email)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("Alefe", "a@a.com", "1")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("1"))
	assert.False(t, user.ValidatePassword("12345"))
	assert.NotEqual(t, "1234", user.Password)
}
