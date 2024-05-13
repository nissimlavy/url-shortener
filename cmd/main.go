package main

import (
    "log"
    "url-shortener/internal/api"
    "url-shortener/internal/store"
    "url-shortener/internal/urlshortener"
    "url-shortener/pkg/config"
)

func main() {
    cfg := config.LoadConfig()
    
    store, err := store.NewStore(cfg.DatabaseURL)
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    
    us := urlshortener.NewURLShortener(store)
    api.StartServer(cfg.ServerPort, us)
}
