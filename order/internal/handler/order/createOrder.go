package order

import (
	"context"
	"github.com/Muvi7z/boilerplate/order/internal/converter"
	generated "github.com/Muvi7z/boilerplate/shared/pkg/server"
)

func (h *Handler) CreateOrder(ctx context.Context, request generated.PostApiV1OrdersRequestObject) (generated.PostApiV1OrdersResponseObject, error) {
	if request.Body == nil {
		return generated.PostApiV1Orders400Response{}, nil
	}
	createOrder := converter.GeneratedToCreateOrderEntity(generated.PostApiV1OrdersJSONBody(*request.Body))
	order, err := h.orderUseCase.CreateOrder(ctx, createOrder)
	if err != nil {
		return generated.PostApiV1Orders500JSONResponse{
			N5xxJSONResponse: struct {
				Code      *int    `json:"code,omitempty"`
				Message   string  `json:"message"`
				RequestId *string `json:"request_id,omitempty"`
			}{
				Message: "Internal error",
			},
		}, nil
	}

	orderUuid := generated.OrderUuid([]byte(order.OrderUuid))

	return generated.PostApiV1Orders200JSONResponse{
		OrderUuid:  &orderUuid,
		TotalPrice: &order.TotalPrice,
	}, nil
}
