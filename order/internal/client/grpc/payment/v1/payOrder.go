package v1

import (
	"context"
	payment_v1 "github.com/Muvi7z/boilerplate/shared/pkg/proto/payment/v1"
)

func (c *client) PayOrder(ctx context.Context, client payment_v1.PaymentClient) (string, error) {
	order, err := client.PayOrder(ctx, &payment_v1.PayOrderRequest{
		OrderUuid:     "",
		UserUuid:      "",
		PaymentMethod: 0,
	})
	if err != nil {
		return "", err
	}

	return order.TransactionUuid, err
}
