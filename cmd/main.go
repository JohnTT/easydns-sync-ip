package main

import (
	"log"
	"time"

	"github.com/JohnTT/easydns-sync-ip/internal/config"
	"github.com/JohnTT/easydns-sync-ip/internal/easydns"
)

var client *easydns.Client

func init() {
	client = easydns.NewClient()
}

func main() {
	// Initial update on program startup.
	log.Printf("client update interval is %v seconds", config.Get().UpdateInterval)
	if err := client.Update(); err != nil {
		log.Printf("client.Update error %v", err)
	}

	// Start periodic sync
	ticker := time.NewTicker(time.Duration(config.Get().UpdateInterval) * time.Second)
	defer ticker.Stop()
	for range ticker.C {
		if err := client.Update(); err != nil {
			log.Printf("client.Update error %v", err)
		}
	}
}
