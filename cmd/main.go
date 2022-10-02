package main

import (
	"database/sql"
	"fmt"

	"github.com/dyhalmeida/fullcycle-golang-intensive/internal/order/entity"
	"github.com/dyhalmeida/fullcycle-golang-intensive/internal/order/infra/database"
)

func main() {
	order, errorNewOrder := entity.NewOrder("1", 30.0, 2.0)

	if errorNewOrder != nil {
		panic(errorNewOrder)
	}

	errorFinalPrice := order.CalculateFinalPrice()

	if errorFinalPrice != nil {
		panic(errorFinalPrice)
	}

	fmt.Printf("The final price is: %f", order.FinalPrice)

	db, errorOpenDb := sql.Open("mysql", "root:root@tcp(mysql:3306)/orders")
	if errorOpenDb != nil {
		panic(errorOpenDb)
	}

	repository := database.NewOrderRepository(db)
	errorSaveOrder := repository.Save(order)
	if errorSaveOrder != nil {
		panic(errorSaveOrder)
	}

}
