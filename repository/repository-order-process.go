package repository

import (
	"database/sql"
	"lumoshive-golang-day-18/model"
)

type OrderProcessRepository interface {
	Create(orderProcess *model.OrderProcess) error
}

type OrderProcessRepositoryDB struct {
	DB *sql.DB
}

func NewOrderProcessRepository(db *sql.DB) OrderProcessRepository {
	return &OrderProcessRepositoryDB{DB: db}
}

func (r *OrderProcessRepositoryDB) Create(orderProcess *model.OrderProcess) error {
	query := `INSERT INTO Order_Process (driver_id, order_id, status) VALUES ($1, $2, $3) RETURNING id`
	err := r.DB.QueryRow(query, orderProcess.DriverID, orderProcess.OrderID, orderProcess.Status).Scan(&orderProcess.ID)
	if err != nil {
		return err
	}

	return nil
}
