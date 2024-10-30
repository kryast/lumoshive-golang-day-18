package repository

import (
	"database/sql"
	"lumoshive-golang-day-18/model"
)

type SessionRepository interface {
	Create(session *model.Session) error
	GetCustomerSession() ([]model.CustomerSession, error)
}

type SessionRepositoryDB struct {
	DB *sql.DB
}

func NewSessionRepository(db *sql.DB) SessionRepository {
	return &SessionRepositoryDB{DB: db}
}

func (r *SessionRepositoryDB) Create(session *model.Session) error {
	query := `INSERT INTO session (customer_id, is_active) VALUES ($1, $2) RETURNING id`
	err := r.DB.QueryRow(query, session.CustomerID, session.IsActive).Scan(&session.ID)
	if err != nil {
		return err
	}

	return nil
}
