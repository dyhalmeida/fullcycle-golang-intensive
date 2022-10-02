package main

import (
	"database/sql"

	"github.com/dyhalmeida/fullcycle-golang-intensive/internal/order/infra/database"
	"github.com/dyhalmeida/fullcycle-golang-intensive/internal/order/usecase"
)

func main() {
	db, errorOpenDb := sql.Open("mysql", "root:root@tcp(mysql:3306)/orders")
	if errorOpenDb != nil {
		panic(errorOpenDb)
	}

	defer db.Close()
	repository := database.NewOrderRepository(db)
	uc := usecase.NewCalculateFinalPriceUseCase(repository)

	input := usecase.OrderInputDTO{
		ID:    "2",
		Price: 56,
		Tax:   7,
	}

	output, errorExecute := uc.Execute(input)

	if errorExecute != nil {
		panic(errorExecute)
	}

	println(output.FinalPrice)

}
