package api

import (
	"net/http"
	"url-shortener/internal/urlshortener"

	"github.com/gin-gonic/gin"
)

// handleShortenURL handles requests to shorten a URL
func handleShortenURL(c *gin.Context, us *urlshortener.URLShortener) {
	var req struct {
		URL string `json:"url"` // Struct to parse incoming JSON with the URL to be shortened
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	shortURL, err := us.ShortenURL(req.URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to shorten URL", "err": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"short_url": shortURL})
}

// handleExpandURL handles redirection from a short URL to the original URL
func handleExpandURL(c *gin.Context, us *urlshortener.URLShortener) {
	shortURL := c.Param("shortUrl")
	originalURL, err := us.ExpandURL(shortURL)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}
	c.Redirect(http.StatusMovedPermanently, originalURL)
}
