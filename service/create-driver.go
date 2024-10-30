package service

import (
	"database/sql"
	"fmt"
	"lumoshive-golang-day-18/model"
	"lumoshive-golang-day-18/repository"
)

func CreateDataDriver(db *sql.DB, name, bike, policeNumber string) error {
	if name == "" {
		return fmt.Errorf("name tidak boleh kosong")
	}
	if bike == "" {
		return fmt.Errorf("bike tidak boleh kosong")
	}
	if policeNumber == "" {
		return fmt.Errorf("police_number tidak boleh kosong")
	}

	driverRepo := repository.NewDriverRepository(db)
	driver := model.Driver{
		Name:         name,
		Bike:         bike,
		PoliceNumber: policeNumber,
	}

	err := driverRepo.Create(&driver)
	if err != nil {
		return fmt.Errorf("gagal input data driver: %w", err)
	}

	fmt.Println("berhasil input data driver dengan id", driver.ID)
	return nil
}
