package main

import (
	"context"
	"log"
	"time"

	"github.com/JohnTT/easydns-sync-ip/internal/config"
	openapi "github.com/JohnTT/easydns-sync-ip/openapi/client/go"
)

var client *openapi.APIClient

func init() {
	cfg := openapi.NewConfiguration()
	cfg.Debug = true
	client = openapi.NewAPIClient(cfg)
}

func update() {
	log.Printf("Ticker triggered at %v", time.Now())
	_, resp, err := client.ZoneAPI.ListZone(context.TODO(), config.Get().Domain).Execute()
	if err != nil {
		log.Printf("error: %v", err)
	} else {
		log.Printf("resp: %+v", resp)
	}
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
