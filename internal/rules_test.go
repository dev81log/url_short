package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateUniqueCode(t *testing.T) {

	code := GenerateUniqueCode()

	assert.NotNil(t, code, "O código não deve ser nulo")
	assert.Equal(t, 6, len(code), "O código deve ter 6 caracteres")
}

func TestGetURLFromCodeIfNotExpired(t *testing.T) {

	url, err := GetURLFromCodeIfNotExpired("LrFuDG")

	assert.Nil(t, err, "Não deve retornar erro")
	assert.NotNil(t, url, "A URL não deve ser nula")
	assert.Equal(t, "http://www.google.com", url, "A URL deve ser http://www.google.com")
}
