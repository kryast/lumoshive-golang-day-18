package repository

import (
	"database/sql"
	"lumoshive-golang-day-18/model"
)

type DriverRepository interface {
	Create(driver *model.Driver) error
	GetBestDriver() ([]model.BestDriver, error)
}

type DriverRepositoryDB struct {
	DB *sql.DB
}

func NewDriverRepository(db *sql.DB) DriverRepository {
	return &DriverRepositoryDB{DB: db}
}

func (r *DriverRepositoryDB) Create(driver *model.Driver) error {
	query := `INSERT INTO driver (name, bike, police_number) VALUES ($1, $2, $3) RETURNING id`
	err := r.DB.QueryRow(query, driver.Name, driver.Bike, driver.PoliceNumber).Scan(&driver.ID)
	if err != nil {
		return err
	}

	return nil
}
