package entity_test

import (
	"testing"

	"github.com/dyhalmeida/fullcycle-golang-intensive/internal/order/entity"
	"github.com/stretchr/testify/assert"
)

func TestGivenEmptyId_WhenCreateANewOrder_ThenShouldReceiveAError(t *testing.T) {
	order := entity.Order{}
	assert.Error(t, order.IsValid(), "invalid id")
}

func TestGivenEmptyPrice_WhenCreateANewOrder_ThenShouldReceiveAError(t *testing.T) {
	order := entity.Order{ID: "1"}
	assert.Error(t, order.IsValid(), "invalid price")
}

func TestGivenEmptyTax_WhenCreateANewOrder_ThenShouldReceiveAError(t *testing.T) {
	order := entity.Order{
		ID:    "1",
		Price: 10,
	}
	assert.Error(t, order.IsValid(), "invalid tax")
}

func TestGivenAValidParams_WhenCallNewOrder_ThenShouldReceiveCreateOrderWithAllParams(t *testing.T) {
	order, err := entity.NewOrder("123", 10.0, 5.0)
	assert.NoError(t, err)
	assert.Equal(t, "123", order.ID)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 5.0, order.Tax)
}

func TestGivenAValidParams_WhenCallCalculateFinalPrice_ThenShouldCalculateFinalPriceAndSetItOnFinalPriceProperty(t *testing.T) {
	order, errorNewOrder := entity.NewOrder("123", 10.0, 5.0)
	assert.NoError(t, errorNewOrder)
	errorFinalPrice := order.CalculateFinalPrice()
	assert.NoError(t, errorFinalPrice)
	assert.Equal(t, 10.5, order.FinalPrice)
}
