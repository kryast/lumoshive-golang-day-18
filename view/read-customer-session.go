package view

import (
	"database/sql"
	"fmt"
	"log"
	"lumoshive-golang-day-18/database"
	"lumoshive-golang-day-18/model"
	"lumoshive-golang-day-18/repository"
)

func GetCustomerSession(db *sql.DB) ([]model.CustomerSession, error) {
	sessionRepo := repository.NewSessionRepository(db)
	customerSession, err := sessionRepo.GetCustomerSession()
	if err != nil {
		return nil, fmt.Errorf("error : %w", err)
	}
	return customerSession, nil
}

func ReadCustomerSession() {

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	customerSession, err := GetCustomerSession(db)
	if err != nil {
		log.Fatalf("error : %v", err)
	}

	fmt.Println("Customer Session :")
	for _, status := range customerSession {
		fmt.Printf("Status: %s, Total Customers: %d\n", status.Status, status.TotalCustomers)
	}
}
