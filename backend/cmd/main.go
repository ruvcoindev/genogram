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
    "github.com/joho/godotenv"
    "go.uber.org/zap"
)

func main() {
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found")
    }

    logger, _ := zap.NewProduction()
    defer logger.Sync()

    logger.Info("Starting Genogram System v1.1")

    router := mux.NewRouter()
    
    // Применяем CORS middleware ко всем маршрутам
    router.Use(corsMiddleware)
    
    api := router.PathPrefix("/api/v1").Subrouter()

    // Health check
    api.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]string{
            "status":  "ok",
            "version": "1.1.0",
            "message": "Genogram System is running",
        })
    }).Methods("GET", "OPTIONS")

    // CBT analyze
    api.HandleFunc("/cbt/analyze", handleCBTAnalyze).Methods("POST", "OPTIONS")

    // Symptoms
    api.HandleFunc("/symptoms", handleGetSymptoms).Methods("GET", "OPTIONS")

    // Psychology traumas
    api.HandleFunc("/psychology/traumas", handleGetTraumas).Methods("GET", "OPTIONS")

    // Affirmations
    api.HandleFunc("/affirmations", handleGetAffirmations).Methods("GET", "OPTIONS")

    // Stats
    api.HandleFunc("/stats", handleGetStats).Methods("GET", "OPTIONS")

    // Diary
    api.HandleFunc("/diary/sections", handleGetDiarySections).Methods("GET", "OPTIONS")
    api.HandleFunc("/diary/questions", handleGetDiaryQuestions).Methods("GET", "OPTIONS")
    api.HandleFunc("/diary/entries", handleSaveDiaryEntry).Methods("POST", "OPTIONS")

    // Frontend static files
    router.PathPrefix("/").Handler(http.FileServer(http.Dir("../frontend")))

    port := getEnv("SERVER_PORT", "8080")
    host := getEnv("SERVER_HOST", "0.0.0.0")
    
    srv := &http.Server{
        Handler:      router,
        Addr:         host + ":" + port,
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }

    go func() {
        logger.Info("Server started", zap.String("address", srv.Addr))
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
