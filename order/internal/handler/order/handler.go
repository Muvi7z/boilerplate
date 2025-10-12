package order

import "github.com/Muvi7z/boilerplate/order/internal/usecase/order"

type Handler struct {
	orderUseCase *order.UseCase
}

func NewHandler(orderUseCase *order.UseCase) *Handler {
	return &Handler{
		orderUseCase: orderUseCase,
	}
}
