package order

import (
	"context"
	"github.com/Muvi7z/boilerplate/order/internal/converter"
	generated "github.com/Muvi7z/boilerplate/shared/pkg/server"
)

func (h *Handler) PayOrder(ctx context.Context, request generated.PostApiV1OrdersOrderUuidPayRequestObject) (generated.PostApiV1OrdersOrderUuidPayResponseObject, error) {
	payOrder := converter.GeneratedPayOrderToEntity(generated.PostApiV1OrdersOrderUuidPayJSONBody(*request.Body))

	transactionUuid, err := h.orderUseCase.PayOrder(ctx, payOrder)
	if err != nil {
		return nil, err
	}

	transactionUuidGenerated := generated.TransactionUuid([]byte(transactionUuid))

	return &generated.PostApiV1OrdersOrderUuidPay200JSONResponse{
		TransactionUuid: &transactionUuidGenerated,
	}, nil
}
