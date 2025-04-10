package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/Carlos6464/imersao22/go-gateway/internal/repository"
	"github.com/Carlos6464/imersao22/go-gateway/internal/service"
	"github.com/Carlos6464/imersao22/go-gateway/internal/web/server"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_PORT", "5432"),
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", "postgres"),
		getEnv("DB_NAME", "postgres"),
	)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	accountRepository := *repository.NewAccountRepository(db)
	accountService := service.NewAccountService(accountRepository)

	port := getEnv("APP_PORT", "8080")
	server := server.NewServer(accountService, port)
	server.ConfigRoute()

	if err := server.Start(); err != nil {
		log.Fatal("Error starting server:", err)
	}

	server.Start()

}
