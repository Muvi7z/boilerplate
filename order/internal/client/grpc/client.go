package grpc

import (
	"context"
	payment_v1 "github.com/Muvi7z/boilerplate/shared/pkg/proto/payment/v1"
)

type PaymentClient interface {
	PayOrder(ctx context.Context, client payment_v1.PaymentClient) (string, error)
}
