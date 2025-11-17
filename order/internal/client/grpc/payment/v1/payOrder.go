package v1

import (
	"context"
	"github.com/Muvi7z/boilerplate/order/internal/converter"
	"github.com/Muvi7z/boilerplate/order/internal/entity"
)

func (c *client) PayOrder(ctx context.Context, payOrder entity.PayOrder) (string, error) {

	order, err := c.grpcPaymentClient.PayOrder(ctx, converter.EntityPayOrderToPaymentPayOrder(payOrder))
	if err != nil {
		return "", err
	}

	return order.TransactionUuid, err
}
