package view

import (
	"database/sql"
	"fmt"
	"log"
	"lumoshive-golang-day-18/database"
	"lumoshive-golang-day-18/model"
	"lumoshive-golang-day-18/repository"
)

func GetBestHourOrder(db *sql.DB) ([]model.BestHourOrder, error) {
	ordersRepo := repository.NewOrdersRepository(db)
	bestHour, err := ordersRepo.GetBestHourOrder()
	if err != nil {
		return nil, fmt.Errorf("error : %w", err)
	}
	return bestHour, nil
}

func ReadBestHourOrder() {
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	bestHour, err := GetBestHourOrder(db)
	if err != nil {
		log.Fatalf("error : %v", err)
	}

	fmt.Println("Best-Hour Order :")
	for _, hour := range bestHour {
		fmt.Printf("Hour: %d:00, Total Orders: %d\n", hour.OrderHour, hour.TotalOrders)
	}
}
