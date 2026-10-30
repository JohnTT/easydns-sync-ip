package main

import (
	"log/slog"
	"os"
	"time"

	"github.com/JohnTT/easydns-sync-ip/internal/config"
	"github.com/JohnTT/easydns-sync-ip/internal/easydns"
)

var client *easydns.Client

func init() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))
	client = easydns.NewClient()
}

func main() {
	// Initial update on program startup.
	slog.Info("client update interval", "seconds", config.Get().UpdateInterval)
	if err := client.Update(); err != nil {
		slog.Error("client.Update failed", "error", err)
	}

	// Start periodic sync
	ticker := time.NewTicker(time.Duration(config.Get().UpdateInterval) * time.Second)
	defer ticker.Stop()
	for range ticker.C {
		if err := client.Update(); err != nil {
			slog.Error("client.Update failed", "error", err)
		}
	}
}
