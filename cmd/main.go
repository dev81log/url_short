package main

import (
	"github.com/dev81log/url_short/internal"
)

func main() {
	internal.Init()
	internal.InsertURLAndCode("http://exemplolinkpostandsave.com", "123456")
	internal.GetURLByCode("123456")
}
