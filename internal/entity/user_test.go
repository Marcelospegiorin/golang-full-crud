package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Marcelo Spegiorin", "marcelo.mmspe@hotmail.com", "123marcelo")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "Marcelo Spegiorin", user.Name)
	assert.Equal(t, "marcelo.mmspe@hotmail.com", user.Email)
}

func TestUserValidatePassword(t *testing.T) {
	user, err := NewUser("Marcelo Spegiorin", "marcelo.mmspe@hotmail.com", "123marcelo")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.True(t, user.ValidatePassword("123marcelo"))
	assert.False(t, user.ValidatePassword("123marcel0"))
	assert.NotEqual(t, "123marcelo", user.Password)
}
