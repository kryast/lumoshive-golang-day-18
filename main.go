package main

import (
	"fmt"
	"lumoshive-golang-day-18/view"

	_ "github.com/lib/pq"
)

func main() {

	view.ReadMontlyOrder()
	fmt.Println("===============")
	view.ReadMonthlyCustomer()
	fmt.Println("===============")
	view.ReadBestLocationOrder()
	fmt.Println("===============")
	view.ReadBestHourOrder()
	fmt.Println("===============")
	view.ReadCustomerSession()
	fmt.Println("===============")
	view.ReadBestDriver()
	fmt.Println("===============")
}
