package order

import (
	"context"
	"errors"
	"github.com/Muvi7z/boilerplate/order/internal/converter"
	"github.com/Muvi7z/boilerplate/order/internal/entity"
	generated "github.com/Muvi7z/boilerplate/shared/pkg/server"
)

func (h *Handler) GetOrder(ctx context.Context, request generated.GetApiV1OrdersOrderUuidRequestObject) (generated.GetApiV1OrdersOrderUuidResponseObject, error) {
	if request.OrderUuid.String() == "" {
		return generated.GetApiV1OrdersOrderUuid400Response{}, nil
	}

	order, err := h.orderUseCase.GetOrder(ctx, request.OrderUuid.String())
	if err != nil {
		if errors.Is(err, entity.ErrOrderNotFound) {
			return generated.GetApiV1OrdersOrderUuid404Response{}, nil
		}

		return generated.GetApiV1OrdersOrderUuid500JSONResponse{
			N5xxJSONResponse: struct {
				Code      *int    `json:"code,omitempty"`
				Message   string  `json:"message"`
				RequestId *string `json:"request_id,omitempty"`
			}{
				Message: "not implemented",
			},
		}, nil
	}

	genResp := converter.OrderEntityFromGenerated(order)

	return generated.GetApiV1OrdersOrderUuid200JSONResponse(genResp), nil
}
