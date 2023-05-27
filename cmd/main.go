package main

import (
	"github.com/dev81log/url_short/internal"
)

func main() {
	internal.Init()
	internal.InsertURLAndCode("http://example.com", "abcdef")

}
