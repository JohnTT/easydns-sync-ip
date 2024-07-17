package main

import (
	"log"
	"time"

	"github.com/JohnTT/easydns-sync-ip/internal/config"
)

func update() {
	log.Printf("Ticker triggered at %v", time.Now())
}

func main() {
	// Initial on program startup.
	update()

	// Start periodic sync
	ticker := time.NewTicker(time.Duration(config.Get().TickerSeconds) * time.Second)
	defer ticker.Stop()
	for range ticker.C {
		update()
	}
}
