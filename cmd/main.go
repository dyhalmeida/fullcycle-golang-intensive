package main

import (
	"fmt"

	"github.com/dyhalmeida/fullcycle-golang-intensive/internal/order/entity"
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
}
