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

func TestGetURLByCode(t *testing.T) {

	url := "https://www.example.com"
	code := "ABC123"

	err := InsertURLAndCode(url, code)

	assert.NoError(t, err, "Não deve ocorrer um erro ao inserir a URL e o código")
	assert.NotNil(t, db, "O banco de dados não deve ser nulo")

	url, err = GetURLByCode(code)

	assert.NoError(t, err, "Não deve ocorrer um erro ao buscar a URL pelo código")
	assert.NotNil(t, db, "O banco de dados não deve ser nulo")
	assert.Equal(t, "https://www.example.com", url, "As URLs devem ser iguais")
}
