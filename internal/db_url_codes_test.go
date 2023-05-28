package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertURLAndCode(t *testing.T) {

	url := "https://www.example.com"
	code := "ABC123"

	err := InsertURLAndCode(url, code)

	assert.NoError(t, err, "Não deve ocorrer um erro ao inserir a URL e o código")
	assert.NotNil(t, db, "O banco de dados não deve ser nulo")
}
