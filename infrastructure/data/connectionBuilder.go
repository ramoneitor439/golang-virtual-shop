package data

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func CreatePostgresqlConnection() (*PostgresqlConnection, error) {
	if envErr := godotenv.Load(".env"); envErr != nil {
		return nil, envErr
	}

	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")
	dbUser := os.Getenv("db_user")
	dbPassword := os.Getenv("db_password")
	dbDatabase := os.Getenv("db_database")

	port, err := strconv.ParseUint(dbPort, 10, 64)
	if err != nil {
		return nil, err
	}

	return &PostgresqlConnection{
		host:              dbHost,
		port:              port,
		user:              dbUser,
		password:          dbPassword,
		database:          dbDatabase,
		connection:        nil,
		activeTransaction: nil,
	}, nil
}
