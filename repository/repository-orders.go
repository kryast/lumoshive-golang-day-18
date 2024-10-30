package repository

import (
	"database/sql"
	"lumoshive-golang-day-18/model"
)

type OrdersRepository interface {
	Create(order *model.Orders) error
	GetMontlyOrder() ([]model.MontlyOrder, error)
	GetMonthlyCustomer() ([]model.MonthlyCustomer, error)
	GetBestLocationOrder() ([]model.BestLocationOrder, error)
	GetBestHourOrder() ([]model.BestHourOrder, error)
}

type OrdersRepositoryDB struct {
	DB *sql.DB
}

func NewOrdersRepository(db *sql.DB) OrdersRepository {
	return &OrdersRepositoryDB{DB: db}
}

func (r *OrdersRepositoryDB) Create(order *model.Orders) error {
	query := `INSERT INTO orders (pick_up, destination, price, customer_id) VALUES ($1, $2, $3, $4) RETURNING id`
	err := r.DB.QueryRow(query, order.PickUp, order.Destination, order.Price, order.CustomerID).Scan(&order.ID)
	if err != nil {
		return err
	}

	return nil
}
