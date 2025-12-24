package app

import (
	"context"
	"github.com/Muvi7z/boilerplate/order/internal/client/grpc"
	inventory "github.com/Muvi7z/boilerplate/order/internal/client/grpc/inventory/v1"
	v1 "github.com/Muvi7z/boilerplate/order/internal/client/grpc/payment/v1"
	"github.com/Muvi7z/boilerplate/order/internal/config"
	inventory_v1 "github.com/Muvi7z/boilerplate/shared/pkg/proto/inventory/v1"
	payment_v1 "github.com/Muvi7z/boilerplate/shared/pkg/proto/payment/v1"
	grpc2 "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type App struct {
	paymentService   grpc.PaymentClient
	inventoryService grpc.InventoryClient
}

func New(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(ctx context.Context) error{}

	for _, init := range inits {
		err := init(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) InitPaymentGRPCClient(ctx context.Context) error {
	connPayment, err := grpc2.NewClient(
		config.AppConfig().PaymentGRPC.Address(),
		grpc2.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return err
	}
	paymentClient := payment_v1.NewPaymentClient(connPayment)

	a.paymentService = v1.New(paymentClient)

	return nil
}

func (a *App) InitInventoryGRPCClient(ctx context.Context) error {
	connInventory, err := grpc2.NewClient(
		config.AppConfig().InventoryGRPC.Address(),
		grpc2.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return err
	}

	inventoryClient := inventory_v1.NewInventoryServiceClient(connInventory)

	a.inventoryService = inventory.New(inventoryClient)

	return nil
}

func (a *App) InitContainer(ctx context.Context) error {

	return nil
}
