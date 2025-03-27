package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var DB *pgxpool.Pool

func InitDB() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Get database URL from environment variables
	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	// Create a connection pool
	DB, err = pgxpool.New(context.Background(), databaseUrl)
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}

	fmt.Println("✅ Connected to PostgreSQL")
}
