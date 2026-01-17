package main

import (
	"database/sql"
	"log"
	"rest-api-gin-go/internal/database"
	"rest-api-gin-go/internal/env"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
)

type application struct {
	port      int
	jwtSecret string
	models database.Models
}

func main() {
	// create database connection
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	models := database.NewModels(db)

	app := &application{
		port:      env.GetEnvInt("PORT", 8080),
		jwtSecret: env.GetEnvString("JWT_SECRET", "your_jwt_secret"),
		models:    models,
	}

	if err := app.server(); err != nil {
		log.Fatal(err)
	}
}
