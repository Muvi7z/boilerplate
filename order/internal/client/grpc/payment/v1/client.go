package v1

import payment_v1 "github.com/Muvi7z/boilerplate/shared/pkg/proto/payment/v1"

type client struct {
	grpcPaymentClient payment_v1.PaymentClient
}

func New(grpcPaymentClient payment_v1.PaymentClient) *client {
	return &client{
		grpcPaymentClient: grpcPaymentClient,
	}
}
