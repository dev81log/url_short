package main

import (
	"github.com/dev81log/url_short/internal"
)

func main() {
	internal.Init()
	//internal.InsertURLAndCode("http://expires_at.com", "exp123")
	// internal.GetURLByCode("123456")
	internal.IsExpiredCode("exp123")
}
