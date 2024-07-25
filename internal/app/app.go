package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	httpInterface "tometower/internal/interface/http"
	"tometower/internal/middleware"
)

type App struct {
	AppRouter *middleware.Logger
}

func NewApp() *App {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	db_url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := sql.Open("postgres", db_url)
	if err != nil {
		log.Fatal(err)
	}

	appRouter := httpInterface.Router(db)

	return &App{
		AppRouter: appRouter,
	}
}

func (app *App) Run(addr string) error {
	log.Printf("Server started on %s", addr)
	return http.ListenAndServe(addr, app.AppRouter)
}
