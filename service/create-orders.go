package service

import (
	"database/sql"
	"fmt"
	"lumoshive-golang-day-18/model"
	"lumoshive-golang-day-18/repository"
)

func CreateDataOrders(db *sql.DB, pickUp, destination string, price int, customerID int) error {
	if pickUp == "" {
		return fmt.Errorf("pick_up tidak boleh kosong")
	}
	if destination == "" {
		return fmt.Errorf("destination tidak boleh kosong")
	}
	if price <= 0 {
		return fmt.Errorf("price harus valid")
	}
	if customerID <= 0 {
		return fmt.Errorf("customer_id harus valid")
	}

	orderRepo := repository.NewOrdersRepository(db)
	order := model.Orders{
		PickUp:      pickUp,
		Destination: destination,
		Price:       price,
		CustomerID:  customerID,
	}

	err := orderRepo.Create(&order)
	if err != nil {
		return fmt.Errorf("gagal input data order: %w", err)
	}

	fmt.Println("berhasil input data order dengan id", order.ID)
	return nil
}
