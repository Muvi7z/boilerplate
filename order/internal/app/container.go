package app

import (
	"context"
	"github.com/Muvi7z/boilerplate/order/internal/client/grpc"
	v1 "github.com/Muvi7z/boilerplate/order/internal/client/grpc/payment/v1"
	"github.com/Muvi7z/boilerplate/order/internal/config"
	"github.com/Muvi7z/boilerplate/order/internal/usecase/order"
	payment_v1 "github.com/Muvi7z/boilerplate/shared/pkg/proto/payment/v1"
	grpc2 "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type container struct {
	paymentService  grpc.PaymentClient
	inventoryClient grpc.InventoryClient

	grpcPaymentClient payment_v1.PaymentClient


	orderRepository order.Repository
}

func NewContainer() *container {
	return &container{}
}

func (c *container) PaymentService(ctx context.Context) grpc.PaymentClient {
	if c.paymentService == nil {
		connPayment, err := grpc2.NewClient(
			config.AppConfig().PaymentGRPC.Address(),
			grpc2.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			return err
		}
		paymentClient := payment_v1.NewPaymentClient(connPayment)

		c.paymentService = v1.New(paymentClient)
	}

	return c.paymentService
}

func (c *container) PaymentClient(ctx context.Context)   {
	if c.grpcPaymentClient == nil {

		c.grpcPaymentClient = payment_v1.NewPaymentClient()
	}
}

func (c *container) InventoryClient(ctx context.Context) payment_v1.PaymentClient {
	if c.inventoryClient == nil {
		c.inventoryClient =
	}
}
