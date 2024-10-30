package view

import (
	"database/sql"
	"fmt"
	"log"
	"lumoshive-golang-day-18/database"
	"lumoshive-golang-day-18/model"
	"lumoshive-golang-day-18/repository"
)

func GetBestDriver(db *sql.DB) ([]model.BestDriver, error) {
	driverRepo := repository.NewDriverRepository(db)
	bestDriver, err := driverRepo.GetBestDriver()
	if err != nil {
		return nil, fmt.Errorf("error : %w", err)
	}
	return bestDriver, nil
}

func ReadBestDriver() {
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	bestDriver, err := GetBestDriver(db)
	if err != nil {
		log.Fatalf("error : %v", err)
	}

	fmt.Println("Best Drivers:")
	for _, driver := range bestDriver {
		fmt.Printf("Driver: %s, Total Orders: %d\n", driver.DriverName, driver.TotalOrders)
	}
}
