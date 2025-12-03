package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"onepace-library/internal/api"
	"onepace-library/internal/config"
	"onepace-library/internal/library"
	"onepace-library/internal/metadata"
)

func main() {
	// ---- Load Configuration ----
	cfg, err := config.Load("data/config.yml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// ---- Load Library JSON ----
	lib, err := library.Load("data/library.json")
	if err != nil {
		log.Printf("No library.json found, creating new empty library...")
		lib = library.New()
	}

	// ---- Create Metadata Client ----
	metaClient := metadata.NewClient(cfg.Metadata.EpisodesURL, cfg.Metadata.ArcsURL)

	if err := metaClient.Refresh(); err != nil {
		log.Fatalf("Failed to load metadata: %v", err)
	}

	// ---- Setup Router ----
	r := chi.NewRouter()

	// Logger, recovery, timeout
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(30 * time.Second))

	// CORS for SvelteKit
	r.Use(api.CORS)

	// ---- Register Routes ----
	api.RegisterRoutes(r, lib, metaClient)

	// ---- Start Server ----
	server := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: r,
	}

	go func() {
		log.Printf("Server started on port %s", cfg.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// ---- Graceful Shutdown ----
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Shutdown failed: %v", err)
	}

	// Save library before quitting
	library.Save("data/library.json", lib)

	log.Println("Server exited gracefully.")
}
