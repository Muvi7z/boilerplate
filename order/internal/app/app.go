package app

import (
	"context"
	"github.com/Muvi7z/boilerplate/order/internal/client/grpc"
	inventory "github.com/Muvi7z/boilerplate/order/internal/client/grpc/inventory/v1"
	v1 "github.com/Muvi7z/boilerplate/order/internal/client/grpc/payment/v1"
	"github.com/Muvi7z/boilerplate/order/internal/config"
	orderhandler "github.com/Muvi7z/boilerplate/order/internal/handler/order"
	"github.com/Muvi7z/boilerplate/order/internal/repository"
	"github.com/Muvi7z/boilerplate/order/internal/server"
	"github.com/Muvi7z/boilerplate/order/internal/usecase/order"
	"github.com/Muvi7z/boilerplate/platform/closer"
	"github.com/Muvi7z/boilerplate/platform/logger"
	inventory_v1 "github.com/Muvi7z/boilerplate/shared/pkg/proto/inventory/v1"
	payment_v1 "github.com/Muvi7z/boilerplate/shared/pkg/proto/payment/v1"
	"github.com/jmoiron/sqlx"
	grpc2 "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type App struct {
	paymentService   grpc.PaymentClient
	inventoryService grpc.InventoryClient

	orderServer *server.Server

	orderHandler *orderhandler.Handler

	orderUsecase *order.UseCase

	db              *sqlx.DB
	orderRepository order.Repository
}

func New(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run(ctx context.Context) error {
	return a.RunOrderServer(ctx)
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(ctx context.Context) error{
		a.initLogger,
		a.initCloser,
		a.initOrderServer,
	}

	for _, init := range inits {
		err := init(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initOrderServer(ctx context.Context) error {
	if a.orderServer == nil {
		orderHandler, err := a.GetOrderHandler(ctx)
		if err != nil {
			return err
		}

		addr := config.AppConfig().AppServerConfig.Address()

		a.orderServer = server.NewServer(orderHandler, addr)
	}

	return nil
}

func (a *App) RunOrderServer(ctx context.Context) error {
	a.orderServer.Run()

	closer.AddNamed("app server", func(ctx context.Context) error {
		err := a.orderServer.Shutdown(ctx)
		if err != nil {
			return err
		}

		return nil
	})

	return nil
}

func (a *App) GetOrderHandler(ctx context.Context) (*orderhandler.Handler, error) {
	if a.orderHandler == nil {
		orderUsecase, err := a.GetOrderUsecase(ctx)
		if err != nil {
			return nil, err
		}
		a.orderHandler = orderhandler.NewHandler(orderUsecase)
	}

	return a.orderHandler, nil
}

func (a *App) GetPaymentGRPCClient(ctx context.Context) (grpc.PaymentClient, error) {
	connPayment, err := grpc2.NewClient(
		config.AppConfig().PaymentGRPC.Address(),
		grpc2.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}
	paymentClient := payment_v1.NewPaymentClient(connPayment)

	a.paymentService = v1.New(paymentClient)

	return a.paymentService, nil
}

func (a *App) GetInventoryGRPCClient(ctx context.Context) (grpc.InventoryClient, error) {
	connInventory, err := grpc2.NewClient(
		config.AppConfig().InventoryGRPC.Address(),
		grpc2.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	inventoryClient := inventory_v1.NewInventoryServiceClient(connInventory)

	a.inventoryService = inventory.New(inventoryClient)

	return a.inventoryService, nil
}

func (a *App) GetOrderRepository(ctx context.Context) (order.Repository, error) {
	if a.orderRepository == nil {
		db, err := sqlx.Connect("postgres", config.AppConfig().Postgres.URI())
		if err != nil {

		}
		a.db = db
		a.orderRepository = repository.New(db)
	}

	return a.orderRepository, nil
}

func (a *App) GetOrderUsecase(ctx context.Context) (*order.UseCase, error) {
	if a.orderUsecase == nil {
		paymentService, err := a.GetPaymentGRPCClient(ctx)
		if err != nil {
			return nil, err
		}
		inventoryService, err := a.GetInventoryGRPCClient(ctx)
		if err != nil {
			return nil, err
		}

		orderRepository, err := a.GetOrderRepository(ctx)
		if err != nil {
			return nil, err
		}
		a.orderUsecase = order.New(paymentService, inventoryService, orderRepository)
	}

	return a.orderUsecase, nil
}

func (a *App) initLogger(ctx context.Context) error {
	return logger.Init(config.AppConfig().Logger.Level(), config.AppConfig().Logger.AsJson())
}

func (a *App) initCloser(_ context.Context) error {
	closer.SetLogger(logger.Logger())
	return nil
}
