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
