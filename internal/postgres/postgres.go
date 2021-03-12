package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/pmaterer/meta/config"
)

func NewHandler(config config.Config) (*sql.DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=%s",
		config.DatabaseHost, config.DatabasePort, config.DatabaseUser,
		config.DatabasePassword, config.DatabaseName, config.DatabaseSSLMode)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return db, err
	}
	return db, nil
}

func Ping(db *sql.DB) error {
	err := db.Ping()
	if err != nil {
		return err
	}
	return nil
}
