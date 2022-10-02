package usecase

import (
	"github.com/dyhalmeida/fullcycle-golang-intensive/internal/order/entity"
)

type OrderInputDTO struct {
	ID    string
	Price float64
	Tax   float64
}

type OrderOutputDTO struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

type CalculateFinalPriceUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func (c *CalculateFinalPriceUseCase) Execute(input OrderInputDTO) (*OrderOutputDTO, error) {

	order, errorNewOrder := entity.NewOrder(input.ID, input.Price, input.Tax)
	if errorNewOrder != nil {
		return nil, errorNewOrder
	}
	errorCalculatePrice := order.CalculateFinalPrice()

	if errorCalculatePrice != nil {
		return nil, errorCalculatePrice
	}

	errorSave := c.OrderRepository.Save(order)
	if errorSave != nil {
		return nil, errorSave
	}

	return &OrderOutputDTO{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}, nil
}

func NewCalculateFinalPriceUseCase(orderRepository entity.OrderRepositoryInterface) *CalculateFinalPriceUseCase {
	return &CalculateFinalPriceUseCase{OrderRepository: orderRepository}
}
