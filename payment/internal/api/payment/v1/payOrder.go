package v1

import (
	"context"
	"github.com/Muvi7z/boilerplate/payment/internal/converter"
	payment_v1 "github.com/Muvi7z/boilerplate/shared/pkg/proto/payment/v1"
)

func (a *api) PayOrder(ctx context.Context, request *payment_v1.PayOrderRequest) (*payment_v1.PayOrderResponse, error) {
	transactionUuid, err := a.paymentService.PayOrder(ctx, converter.PaymentToPaymentEntity(request))
	if err != nil {
		return nil, err
	}

	return &payment_v1.PayOrderResponse{
		TransactionUuid: transactionUuid,
	}, nil
}
