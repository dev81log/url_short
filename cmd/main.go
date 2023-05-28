package main

import (
	"net/http"

	"github.com/dev81log/url_short/internal"
	"github.com/gin-gonic/gin"
)

type requestPayload struct {
	URL string `json:"url" binding:"required"`
}

func shortenURL(c *gin.Context) {
	var payload requestPayload

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	code := internal.GenerateUniqueCode()

	err := internal.InsertURLAndCode(payload.URL, code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not insert URL and code into database."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": code})
}

func redirectURL(c *gin.Context) {
	code := c.Param("code")

	url, err := internal.GetURLFromCodeIfNotExpired(code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusMovedPermanently, url)
}

func main() {
	internal.Init()

	r := gin.Default()

	r.POST("/", shortenURL)
	r.GET("/:code", redirectURL)

	r.Run()
}
