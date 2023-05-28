package internal

import (
	"errors"
	"math/rand"
	"time"
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
