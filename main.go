package main

import (
	"context"
	"log"
	"matrixor/config"
	"matrixor/db"
	"matrixor/server"
)

func main() {
	cfg := config.Load()

	client, err := db.Init(cfg.Database)
	if err != nil {
		log.Fatalf("Failed to initialize DB: %v", err)
	}
	defer client.Disconnect(context.Background())

	srv := server.New(cfg.Server, client, cfg.Database)
	srv.Start()
}
