package service

import (
	"database/sql"
	"errors"
	"fmt"
	"lumoshive-golang-day-18/model"
	"lumoshive-golang-day-18/repository"
)

func CreateDataCustomer(db *sql.DB, name, email string, phone int) error {
	if name == "" {
		return errors.New("name tidak boleh kosong")
	}
	if email == "" {
		return errors.New("email tidak boleh kosong")
	}
	if phone <= 0 {
		return errors.New("phone harus valid")
	}

	customerRepo := repository.NewCustomerRepository(db)
	customer := model.Customer{
		Name:  name,
		Email: email,
		Phone: phone,
	}

	err := customerRepo.Create(&customer)
	if err != nil {
		return fmt.Errorf("gagal input data customer: %w", err)
	}

	fmt.Println("berhasil input data customer dengan id", customer.ID)

	return nil
}
