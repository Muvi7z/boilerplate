package order

import (
	"context"
	"github.com/Muvi7z/boilerplate/order/internal/converter"
	generated "github.com/Muvi7z/boilerplate/shared/pkg/server"
	"github.com/google/uuid"
)

func (h *Handler) PayOrder(ctx context.Context, request generated.PostApiV1OrdersOrderUuidPayRequestObject) (generated.PostApiV1OrdersOrderUuidPayResponseObject, error) {
	payOrder := converter.GeneratedPayOrderToEntity(generated.PostApiV1OrdersOrderUuidPayJSONBody(*request.Body))
	payOrder.OrderUuid = request.OrderUuid.String()

	transactionUuid, err := h.orderUseCase.PayOrder(ctx, payOrder)
	if err != nil {
		return generated.PostApiV1OrdersOrderUuidPay500JSONResponse{
			N5xxJSONResponse: struct {
				Code      *int    `json:"code,omitempty"`
				Message   string  `json:"message"`
				RequestId *string `json:"request_id,omitempty"`
			}{
				Message: "Internal error",
			},
		}, nil
	}

	var transactionUuidGenerated generated.TransactionUuid
	if transactionUuid != "" {
		transactionUuidGenerated, err = uuid.Parse(transactionUuid)
		if err != nil {
			return generated.PostApiV1OrdersOrderUuidPay500JSONResponse{
				N5xxJSONResponse: struct {
					Code      *int    `json:"code,omitempty"`
					Message   string  `json:"message"`
					RequestId *string `json:"request_id,omitempty"`
				}{
					Message: "Internal error",
				},
			}, nil
		}
	}
	return &generated.PostApiV1OrdersOrderUuidPay200JSONResponse{
		TransactionUuid: &transactionUuidGenerated,
	}, nil
}
