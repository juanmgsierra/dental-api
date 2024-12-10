package main

import (
	"context"
	"juanmgsierra/dental-api/api"
	db "juanmgsierra/dental-api/db/sqlc"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	connStr := os.Getenv("DATABASE_URL")

	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		log.Fatal("Unable to parse database config:", err)
	}

	config.MaxConns = 10
	config.MaxConnLifetime = 5 * time.Minute
	config.HealthCheckPeriod = 1 * time.Minute

	// Crear el pool de conexiones
	dbPool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Fatal("Unable to create connection pool:", err)
	}
	defer dbPool.Close()

	// Crear queries usando el pool en lugar de una conexión
	queries := db.New(dbPool)

	server := api.NewServer(queries)

	serverAddress := os.Getenv("SERVER_ADDRESS")
	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
