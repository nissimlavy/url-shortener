package urlshortener

import (
	"crypto/sha256"
	"encoding/base64"
	"log"
	"strings"

	"url-shortener/internal/store"
)

type URLShortener struct {
	store *store.Store
}

func NewURLShortener(s *store.Store) *URLShortener {
	return &URLShortener{store: s}
}

func (us *URLShortener) ShortenURL(originalURL string) (string, error) {
	// Generate a short URL
	log.Print("here")
	sum := sha256.Sum256([]byte(originalURL))
	log.Print("here", sum)
	shortURL := base64.URLEncoding.EncodeToString(sum[:])[:8] // Shorten to 8 characters
	log.Print("here", shortURL)
	shortURL = strings.ReplaceAll(shortURL, "/", "_")
	log.Print("here", shortURL)

	// Store in DB
	err := us.store.CreateURLMapping(shortURL, originalURL)
	if err != nil {
		return "", err
	}
	return shortURL, nil
}

func (us *URLShortener) ExpandURL(shortURL string) (string, error) {
	return us.store.GetOriginalURL(shortURL)
}
