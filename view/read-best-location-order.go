package view

import (
	"database/sql"
	"fmt"
	"log"
	"lumoshive-golang-day-18/database"
	"lumoshive-golang-day-18/model"
	"lumoshive-golang-day-18/repository"
)

func GetBestLocationOrder(db *sql.DB) ([]model.BestLocationOrder, error) {
	ordersRepo := repository.NewOrdersRepository(db)
	bestLocation, err := ordersRepo.GetBestLocationOrder()
	if err != nil {
		return nil, fmt.Errorf("error : %w", err)
	}
	return bestLocation, nil
}

func ReadBestLocationOrder() {
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	bestLocation, err := GetBestLocationOrder(db)
	if err != nil {
		log.Fatalf("error : %v", err)
	}

	fmt.Println("Best Location with Highest Total Orders:")
	for _, location := range bestLocation {
		fmt.Printf("Location: %s, Total Orders: %d\n", location.Location, location.TotalOrders)
	}
}
