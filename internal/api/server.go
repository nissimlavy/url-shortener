package api

import (
    "github.com/gin-gonic/gin"
    "url-shortener/internal/urlshortener"
)

// StartServer initializes and starts the HTTP server
func StartServer(port string, us *urlshortener.URLShortener) {
    router := gin.Default()
    router.POST("/shorten", func(c *gin.Context) {
        handleShortenURL(c, us)
    })
    router.GET("/:shortUrl", func(c *gin.Context) {
        handleExpandURL(c, us)
    })
    router.Run(":" + port)
}
