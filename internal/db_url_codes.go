package internal

import (
	"log"
)

func InsertURLAndCode(url string, code string) error {
	var urlID int
	err := db.QueryRow("INSERT INTO urls (url, created_at, expires_at) VALUES ($1, NOW(), NOW() + INTERVAL '7 days') RETURNING id", url).Scan(&urlID)
	if err != nil {
		return err
	}
	log.Printf("URL Post: %s", url)

	_, err = db.Exec("INSERT INTO url_codes (url_id, code, created_at) VALUES ($1, $2, NOW())", urlID, code)
	log.Printf("Code Post: %s", code)

	return err
}

func GetURLByCode(code string) (string, error) {
	var url string
	err := db.QueryRow("SELECT u.url FROM urls u JOIN url_codes c ON c.url_id = u.id WHERE c.code = $1", code).Scan(&url)
	if err != nil {
		return "", err
	}
	log.Printf("URL Get: %s", url)

	return url, nil

}

func IsExpiredCode(code string) (bool, error) {
	var expired bool
	err := db.QueryRow("SELECT u.expires_at < NOW() FROM urls u JOIN url_codes c ON c.url_id = u.id WHERE c.code = $1", code).Scan(&expired)
	if err != nil {
		return false, err
	}
	log.Printf("Expired: %t", expired)

	return expired, nil
}
