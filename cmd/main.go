package main

import (
	"github.com/dev81log/url_short/internal"
	"github.com/gin-gonic/gin"
)

func main() {
	internal.Init()

	r := gin.Default()

	r.POST("/", internal.ShortenURL)
	r.GET("/:code", internal.RedirectURL)

	r.Run()
}
