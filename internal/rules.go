package internal

import (
	"errors"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateUniqueCode() string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, 6)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func GetURLFromCodeIfNotExpired(code string) (string, error) {
	expired, err := IsExpiredCode(code)
	if err != nil {
		return "", err
	}
	if expired {
		return "", errors.New("The provided code is expired.")
	}

	url, err := GetURLByCode(code)
	if err != nil {
		return "", err
	}

	return url, nil
}

type requestPayload struct {
	URL string `json:"url" binding:"required"`
}

func ShortenURL(c *gin.Context) {
	var payload requestPayload

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	code := GenerateUniqueCode()

	err := InsertURLAndCode(payload.URL, code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not insert URL and code into database."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": code})
}

func RedirectURL(c *gin.Context) {
	code := c.Param("code")

	url, err := GetURLFromCodeIfNotExpired(code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusMovedPermanently, url)
}
