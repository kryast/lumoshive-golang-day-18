package view

import (
	"database/sql"
	"fmt"
	"log"
	"lumoshive-golang-day-18/database"
	"lumoshive-golang-day-18/model"
	"lumoshive-golang-day-18/repository"
)

func GetMontlyOrder(db *sql.DB) ([]model.MontlyOrder, error) {
	ordersRepo := repository.NewOrdersRepository(db)
	monthlyOrder, err := ordersRepo.GetMontlyOrder()
	if err != nil {
		return nil, fmt.Errorf("error : %w", err)
	}
	return monthlyOrder, nil
}

func ReadMontlyOrder() {
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	monthlyOrder, err := GetMontlyOrder(db)
	if err != nil {
		log.Fatalf("error : %v", err)
	}

	fmt.Println("Total Orders by Month:")
	for _, order := range monthlyOrder {
		fmt.Printf("Month: %s, Total Orders: %d\n", order.Month, order.TotalOrder)
	}
}
