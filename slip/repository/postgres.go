package repository

import (
	"database/sql"

	"github.com/lib/pq"
	"github.com/pmaterer/meta/slip"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) CreateSlip(slip slip.Slip) error {
	query := `INSERT INTO slips(body, tags) VALUES($1, $2)`
	statement, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = statement.Exec(slip.Body, pq.Array(slip.Tags))
	if err != nil {
		return err
	}
	return nil
}
