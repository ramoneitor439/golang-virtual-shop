package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"mystore.com/domain/constants/environments"
	"mystore.com/infrastructure"
	"mystore.com/infrastructure/data"
)

func main() {
	log.Println("Starting server app...")

	log.Println("Loading environment variables...")
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error trying to load environment variables. Message: %s \n", err.Error())
		return
	}

	if os.Getenv("go_environment") == environments.DEVELOPMENT {
		migrateDatabase()
	}

	startServer()
}

func startServer() {

	mux := mux.NewRouter()

	// Auth
	mapAuthController(mux)

	log.Println("Server running at http://localhost:8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Error starting server: %s \n", err.Error())
	}
}

func migrateDatabase() {
	connection, connectionErr := data.CreatePostgresqlConnection()
	if connectionErr != nil {
		log.Fatalf("Error connecting to database. Message: %s \n", connectionErr.Error())
		return
	}

	log.Println("Applying migrations...")

	if migrationsErr := infrastructure.ApplyMigrations(connection); migrationsErr != nil {
		log.Fatalf("Error running migrations to database. Message: %s \n", migrationsErr.Error())
		return
	}

	fmt.Println("Migrations applied")
}
