package main

import (
	"blog-api/config"
	"blog-api/internal/database"
	"blog-api/internal/server"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	repo, err := database.NewRepo(ctx, cfg.Mongo.URI, cfg.Mongo.DBName, cfg.Mongo.CollName)
	if err != nil {
		log.Fatal(err)
	}
	server := server.NewServer(cfg.App.Port, repo)

	go func() {
		defer func() {
			r := recover()
			if r != nil {
				log.Printf("recovered from panic: %s", r)
			}
		}()
		err := server.Run()
		if err != nil {
			log.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan

	graceCtx, graceCancel := context.WithTimeout(ctx, 10*time.Second)
	defer graceCancel()

	err = server.Close(graceCtx)
	if err != nil {
		log.Printf("error while closing server: %v", err)
	}
	err = repo.Close()
	if err != nil {
		log.Printf("error while closing mongo client: %v", err)
	}

	log.Println("app was closed successfully")
}
