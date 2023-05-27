package internal

func InsertURLAndCode(url string, code string) error {
	var urlID int
	err := db.QueryRow("INSERT INTO urls (url, created_at, expires_at) VALUES ($1, NOW(), NOW() + INTERVAL '7 days') RETURNING id", url).Scan(&urlID)
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO url_codes (url_id, code, created_at) VALUES ($1, $2, NOW())", urlID, code)
	return err
}
