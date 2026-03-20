package main

import (
    "context"
    "encoding/json"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"

    "github.com/gorilla/mux"
    "github.com/gorilla/websocket"
    "github.com/joho/godotenv"
    "go.uber.org/zap"
)

func main() {
    // Load .env
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found")
    }

    // Logger
    logger, _ := zap.NewProduction()
    defer logger.Sync()

    logger.Info("Starting Genogram System v1.1")

    // Router
    router := mux.NewRouter()

    // API v1
    api := router.PathPrefix("/api/v1").Subrouter()

    // Health check
    api.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]string{
            "status": "ok",
            "version": "1.1.0",
            "message": "Genogram System is running",
        })
    })

    // Serve frontend static files
    router.PathPrefix("/").Handler(http.FileServer(http.Dir("../frontend")))

    // Server
    srv := &http.Server{
        Handler:      router,
        Addr:         ":" + getEnv("SERVER_PORT", "8080"),
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }

    // Graceful shutdown
    go func() {
        logger.Info("Server started", zap.String("port", getEnv("SERVER_PORT", "8080")))
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            logger.Fatal("Server failed", zap.Error(err))
        }
    }()

    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit

    logger.Info("Shutting down server...")

    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    if err := srv.Shutdown(ctx); err != nil {
        logger.Fatal("Server shutdown failed", zap.Error(err))
    }

    logger.Info("Server stopped")
}

func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}
