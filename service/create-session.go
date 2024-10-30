package service

import (
	"database/sql"
	"fmt"
	"lumoshive-golang-day-18/model"
	"lumoshive-golang-day-18/repository"
)

func CreateDataSession(db *sql.DB, customerID int, isActive bool) error {
	if customerID <= 0 {
		return fmt.Errorf("customer_id harus valid")
	}

	sessionRepo := repository.NewSessionRepository(db)
	session := model.Session{
		CustomerID: customerID,
		IsActive:   isActive,
	}

	err := sessionRepo.Create(&session)
	if err != nil {
		return fmt.Errorf("gagal input data session: %w", err)
	}

	fmt.Println("berhasil input data session dengan id", session.ID)
	return nil
}
