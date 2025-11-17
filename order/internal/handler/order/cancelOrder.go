package order

import (
	"context"
	"errors"
	"github.com/Muvi7z/boilerplate/order/internal/entity"
	generated "github.com/Muvi7z/boilerplate/shared/pkg/server"
)

func (h *Handler) CancelOrder(ctx context.Context, request generated.PostApiV1OrdersOrderUuidCancelRequestObject) (generated.PostApiV1OrdersOrderUuidCancelResponseObject, error) {
	if request.OrderUuid.String() != "" {
		return generated.PostApiV1OrdersOrderUuidCancel404Response{}, nil
	}

	err := h.orderUseCase.CancelOrder(ctx, request.OrderUuid.String())
	if err != nil {
		if errors.Is(err, entity.ErrOrderIsPaid) {
			return generated.PostApiV1OrdersOrderUuidCancel409Response{}, nil
		}

		if errors.Is(err, entity.ErrOrderNotFound) {
			return generated.PostApiV1OrdersOrderUuidCancel404Response{}, nil
		}
		return generated.PostApiV1OrdersOrderUuidCancel500JSONResponse{
			N5xxJSONResponse: struct {
				Code      *int    `json:"code,omitempty"`
				Message   string  `json:"message"`
				RequestId *string `json:"request_id,omitempty"`
			}{
				Message: "not implemented",
			},
		}, nil
	}

	return generated.PostApiV1OrdersOrderUuidCancel204Response{}, nil
}
