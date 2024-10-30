package repository

import (
	"database/sql"
	"lumoshive-golang-day-18/model"
)

type CustomerRepository interface {
	Create(customer *model.Customer) error
}

type CustomerRepositoryDB struct {
	DB *sql.DB
}

func NewCustomerRepository(db *sql.DB) CustomerRepository {
	return &CustomerRepositoryDB{DB: db}
}

func (r *CustomerRepositoryDB) Create(customer *model.Customer) error {
	query := `INSERT INTO customer (name, email, phone) VALUES ($1, $2, $3) RETURNING id`
	err := r.DB.QueryRow(query, customer.Name, customer.Email, customer.Phone).Scan(&customer.ID)
	if err != nil {
		return err
	}

	return nil
}
