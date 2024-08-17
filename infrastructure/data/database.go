package data

import "database/sql"

type DataBaseConnection interface {
	Connect() (*sql.DB, error)
	Query(query string, params ...any)
	Mute(sentence string, params ...any)
	Close() error
	StartTransaction() error
	RollbackTransaction() error
	CommitTransaction() error
}
