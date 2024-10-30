package view

import (
	"database/sql"
	"fmt"
	"log"
	"lumoshive-golang-day-18/database"
	"lumoshive-golang-day-18/model"
	"lumoshive-golang-day-18/repository"
)

func GetMonthlyCustomer(db *sql.DB) ([]model.MonthlyCustomer, error) {
	ordersRepo := repository.NewOrdersRepository(db)
	monthlyCustomer, err := ordersRepo.GetMonthlyCustomer()
	if err != nil {
		return nil, fmt.Errorf("error : %w", err)
	}
	return monthlyCustomer, nil
}

func ReadMonthlyCustomer() {
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	monthlyCustomer, err := GetMonthlyCustomer(db)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Println("Customer with Highest Total Orders in a Month:")
	for _, result := range monthlyCustomer {
		fmt.Printf("Month: %s, Total Orders: %d, Customer Name: %s\n", result.Month, result.TotalOrder, result.CustomerName)
	}
}
