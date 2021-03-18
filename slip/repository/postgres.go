package repository

import (
	"database/sql"
	"log"

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

func (r *Repository) GetSlip(id int64) (slip.Slip, error) {
	var slip slip.Slip
	err := r.db.QueryRow("SELECT id, body, tags, created_at, updated_at FROM slips WHERE id = $1", id).Scan(&slip.ID, &slip.Body, pq.Array(&slip.Tags), &slip.CreatedAt, &slip.UpdatedAt)
	if err != nil {
		return slip, err
	}
	return slip, nil
}

func (r *Repository) GetAllSlips() ([]slip.Slip, error) {
	var slips []slip.Slip
	rows, err := r.db.Query("SELECT * FROM slips")
	if err != nil {
		return slips, nil
	}
	defer rows.Close()

	for rows.Next() {
		var slip slip.Slip
		err = rows.Scan(&slip.ID, &slip.Body, pq.Array(&slip.Tags), &slip.CreatedAt, &slip.UpdatedAt)
		if err != nil {
			return slips, err
		}
		slips = append(slips, slip)
	}
	err = rows.Err()
	if err != nil {
		return slips, err
	}
	return slips, nil
}

func (r *Repository) UpdateSlip(slip slip.Slip) error {
	query := `UPDATE slips SET body = $1, tags = $2 WHERE id=$3`
	statement, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	log.Println(slip)
	rowID, err := statement.Exec(slip.Body, pq.Array(slip.Tags), slip.ID)
	if err != nil {
		return err
	}
	id, _ := rowID.RowsAffected()
	log.Printf("Row updated: %d", id)
	return nil
}

func (r *Repository) DeleteSlip(id int64) error {
	query := `DELETE FROM slips WHERE id=$1`
	statement, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = statement.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
