package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

type PostgresqlConnection struct {
	DataBaseConnection
	host              string
	port              uint64
	user              string
	password          string
	database          string
	connection        *sql.DB
	activeTransaction *sql.Tx
}

func (db *PostgresqlConnection) Connect() error {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", db.host, db.port, db.user, db.password, db.database)

	conn, connError := sql.Open("postgres", connectionString)
	if connError != nil {
		return connError
	}

	db.connection = conn

	return nil
}

func (db *PostgresqlConnection) Mute(query string, params ...any) (sql.Result, error) {
	return db.connection.Exec(query, params...)
}

func (db *PostgresqlConnection) Query(query string, params ...any) (*sql.Rows, error) {
	return db.connection.Query(query, params...)
}

func (db *PostgresqlConnection) Close() error {
	return db.connection.Close()
}

func (db *PostgresqlConnection) StartTransaction() error {
	if db.activeTransaction != nil {
		return errors.New("one transaction is already running")
	}

	ctx := context.Background()
	transaction, err := db.connection.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	db.activeTransaction = transaction
	return nil
}

func (db *PostgresqlConnection) RollbackTransaction() error {
	if db.activeTransaction == nil {
		return errors.New("there is no any transaction running currently")
	}

	err := db.activeTransaction.Rollback()
	if err != nil {
		return err
	}

	db.activeTransaction = nil

	return nil
}

func (db *PostgresqlConnection) CommitTransaction() error {
	if db.activeTransaction == nil {
		return errors.New("there is no any transaction running currently")
	}

	err := db.activeTransaction.Commit()
	if err != nil {
		return err
	}

	db.activeTransaction = nil

	return nil
}
