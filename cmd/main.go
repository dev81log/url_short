package main

import (
	"github.com/dev81log/url_short/internal"
)

func main() {
	internal.Init()
	internal.InsertURLAndCode("http://laismeuamor.com", "liah12")
	internal.GetURLByCode("liah12")
	internal.IsExpiredCode("liah12")
}
