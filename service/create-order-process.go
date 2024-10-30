package service

import (
	"database/sql"
	"fmt"
	"lumoshive-golang-day-18/model"
	"lumoshive-golang-day-18/repository"
)

func CreateDataOrderProcess(db *sql.DB, driverID, orderID int, status string) error {
	if driverID <= 0 {
		return fmt.Errorf("driver_id harus valid")
	}
	if orderID <= 0 {
		return fmt.Errorf("order_id harus valid")
	}
	if status == "" {
		return fmt.Errorf("status tidak boleh kosong")
	}

	orderProcessRepo := repository.NewOrderProcessRepository(db)
	orderProcess := model.OrderProcess{
		DriverID: driverID,
		OrderID:  orderID,
		Status:   status,
	}

	err := orderProcessRepo.Create(&orderProcess)
	if err != nil {
		return fmt.Errorf("gagal input data order process: %w", err)
	}

	fmt.Println("berhasil input data order process dengan id", orderProcess.ID)
	return nil
}
