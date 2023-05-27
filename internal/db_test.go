package internal

import (
	"bytes"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	// Mock environment variables
	os.Setenv("POSTGRES_USER", "username")
	os.Setenv("POSTGRES_PASSWORD", "password")
	os.Setenv("POSTGRES_DB", "dbname")
	os.Setenv("POSTGRES_HOST", "localhost")
	os.Setenv("POSTGRES_PORT", "5432")

	var logOutput bytes.Buffer
	log.SetOutput(&logOutput)

	Init()

	// Check if log contains success message
	assert.Contains(t, logOutput.String(), "Successfully connected!")

	assert.NotNil(t, db)

	// Reset environment variables
	os.Unsetenv("POSTGRES_USER")
	os.Unsetenv("POSTGRES_PASSWORD")
	os.Unsetenv("POSTGRES_DB")
	os.Unsetenv("POSTGRES_HOST")
	os.Unsetenv("POSTGRES_PORT")
}
