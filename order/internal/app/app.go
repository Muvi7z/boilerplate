package app

import (
	"context"
	"github.com/IBM/sarama"
	"github.com/Muvi7z/boilerplate/order/internal/client/grpc"
	inventory "github.com/Muvi7z/boilerplate/order/internal/client/grpc/inventory/v1"
	v1 "github.com/Muvi7z/boilerplate/order/internal/client/grpc/payment/v1"
	"github.com/Muvi7z/boilerplate/order/internal/config"
	"github.com/Muvi7z/boilerplate/order/internal/converter/kafka"
	"github.com/Muvi7z/boilerplate/order/internal/converter/kafka/decoder"
	orderhandler "github.com/Muvi7z/boilerplate/order/internal/handler/order"
	"github.com/Muvi7z/boilerplate/order/internal/migrator"
	"github.com/Muvi7z/boilerplate/order/internal/repository"
	"github.com/Muvi7z/boilerplate/order/internal/server"
	"github.com/Muvi7z/boilerplate/order/internal/service"
	"github.com/Muvi7z/boilerplate/order/internal/service/consumer/order_consumer"
	"github.com/Muvi7z/boilerplate/order/internal/service/producer/order_producer"
	"github.com/Muvi7z/boilerplate/order/internal/usecase/order"
	"github.com/Muvi7z/boilerplate/platform/closer"
	kafka3 "github.com/Muvi7z/boilerplate/platform/kafka"
	"github.com/Muvi7z/boilerplate/platform/kafka/consumer"
	"github.com/Muvi7z/boilerplate/platform/kafka/producer"
	"github.com/Muvi7z/boilerplate/platform/logger"
	kafka2 "github.com/Muvi7z/boilerplate/platform/middleware/kafka"
	inventory_v1 "github.com/Muvi7z/boilerplate/shared/pkg/proto/inventory/v1"
	payment_v1 "github.com/Muvi7z/boilerplate/shared/pkg/proto/payment/v1"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	grpc2 "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type App struct {
	paymentService   grpc.PaymentClient
	inventoryService grpc.InventoryClient

	orderServer  *server.Server
	orderHandler *orderhandler.Handler
	orderUsecase *order.UseCase

	migrator *migrator.Migrator

	db              *sqlx.DB
	orderRepository order.Repository

	orderConsumerService service.OrderConsumerService
	orderConsumer        kafka3.Consumer
	orderAssemblyDecoder kafka.OrderAssemblyDecoder
	consumerGroup        sarama.ConsumerGroup

	orderProducerService service.OrderProducerService
	orderProducer        kafka3.Producer
	syncProducer         sarama.SyncProducer
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
	errCh := make(chan error, 2)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go func() {
		if err := a.RunOrderConsumerService(ctx); err != nil {
			errCh <- errors.Errorf("consumer crashed: %v", err)
		}
	}()

	go func() {
		if err := a.RunOrderServer(ctx); err != nil {
			errCh <- errors.Errorf("server crashed: %v", err)
		}
	}()

	select {
	case <-ctx.Done():
		logger.Info(ctx, "shutdown signal received")
	case err := <-errCh:
		logger.Error(ctx, "component crashed", zap.Error(err))

		cancel()

		<-ctx.Done()

		return err
	}

	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(ctx context.Context) error{
		a.initLogger,
		a.initCloser,
		a.initOrderServer,
		a.initOrderConsumerService,
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
		closer.AddNamed("Order server", func(ctx context.Context) error {
			serr := a.orderServer.Shutdown(ctx)
			if serr != nil {
				return serr
			}

			return nil
		})
	}

	return nil
}

func (a *App) RunOrderConsumerService(ctx context.Context) error {
	err := a.orderConsumerService.RunConsumer(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) RunOrderServer(ctx context.Context) error {
	a.orderServer.Run(ctx)

	closer.AddNamed("app server", func(ctx context.Context) error {
		err := a.orderServer.Shutdown(ctx)
		if err != nil {
			return err
		}

		return nil
	})

	return nil
}

func (a *App) GetMigrator(ctx context.Context) (*migrator.Migrator, error) {
	if a.migrator == nil {

		db, err := a.GetDatabase(ctx)
		if err != nil {
			return nil, err
		}
		a.migrator = migrator.NewMigrator(db.DB, config.AppConfig().AppServerConfig.MigrationsDir())
	}

	return a.migrator, nil
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

	closer.AddNamed("inventory gRPC client", func(ctx context.Context) error {
		cerr := connPayment.Close()
		if cerr != nil {
			return cerr
		}

		return nil
	})

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

	closer.AddNamed("inventory gRPC client", func(ctx context.Context) error {
		cerr := connInventory.Close()
		if cerr != nil {
			return cerr
		}

		return nil
	})

	inventoryClient := inventory_v1.NewInventoryServiceClient(connInventory)

	a.inventoryService = inventory.New(inventoryClient)

	return a.inventoryService, nil
}

func (a *App) GetDatabase(ctx context.Context) (*sqlx.DB, error) {
	if a.db == nil {
		uri := config.AppConfig().Postgres.URI()
		db, err := sqlx.Connect("postgres", uri)
		if err != nil {
			return nil, err
		}

		closer.AddNamed("Postgresql client", func(ctx context.Context) error {
			return db.Close()
		})

		a.db = db
	}

	return a.db, nil
}

func (a *App) GetOrderRepository(ctx context.Context) (order.Repository, error) {
	if a.orderRepository == nil {

		db, err := a.GetDatabase(ctx)
		if err != nil {
			return nil, err
		}

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

		orderProducerService, err := a.GetOrderProducerService(ctx)

		a.orderUsecase = order.New(paymentService, inventoryService, orderRepository, orderProducerService)
	}

	return a.orderUsecase, nil
}

func (a *App) GetOrderConsumerGroup(ctx context.Context) (sarama.ConsumerGroup, error) {
	if a.consumerGroup == nil {
		consumerGroup, err := sarama.NewConsumerGroup(
			config.AppConfig().Kafka.Brokers(),
			config.AppConfig().OrderAssembledConsumerConfig.GroupId(),
			config.AppConfig().OrderAssembledConsumerConfig.Config(),
		)

		if err != nil {
			return nil, err
		}
		a.consumerGroup = consumerGroup

		closer.AddNamed("consumer group", func(ctx context.Context) error {
			err := consumerGroup.Close()
			if err != nil {
				return err
			}

			return nil
		})
	}

	return a.consumerGroup, nil
}

func (a *App) GetOrderConsumer(ctx context.Context) (kafka3.Consumer, error) {
	if a.orderConsumer == nil {
		consumerGroup, err := a.GetOrderConsumerGroup(ctx)
		if err != nil {
			return nil, err
		}

		a.orderConsumer = consumer.NewConsumer(
			consumerGroup,
			[]string{
				config.AppConfig().OrderAssembledConsumerConfig.Topic(),
			},
			logger.Logger(),
			kafka2.Logging(logger.Logger()),
		)
	}

	return a.orderConsumer, nil
}

func (a *App) initOrderConsumerService(ctx context.Context) error {
	if a.orderConsumerService == nil {
		orderConsumer, err := a.GetOrderConsumer(ctx)
		if err != nil {
			return err
		}

		orderDecoder := a.GetOrderDecoder(ctx)

		orderUseCase, err := a.GetOrderUsecase(ctx)
		if err != nil {
			return err
		}

		a.orderConsumerService = order_consumer.NewService(orderConsumer, orderDecoder, orderUseCase)
	}

	return nil
}

func (a *App) GetOrderDecoder(ctx context.Context) kafka.OrderAssemblyDecoder {
	if a.orderAssemblyDecoder == nil {
		a.orderAssemblyDecoder = decoder.NewOrderDecoder()
	}

	return a.orderAssemblyDecoder
}

func (a *App) GetOrderProducerService(ctx context.Context) (service.OrderProducerService, error) {
	if a.orderProducerService == nil {
		orderProducer, err := a.GetOrderProducer(ctx)
		if err != nil {
			return nil, err
		}

		a.orderProducerService = order_producer.NewService(orderProducer)
	}

	return a.orderProducerService, nil
}

func (a *App) GetSyncProducer(ctx context.Context) (sarama.SyncProducer, error) {
	if a.syncProducer == nil {
		p, err := sarama.NewSyncProducer(
			config.AppConfig().Kafka.Brokers(),
			config.AppConfig().OrderPaidProducerConfig.Config(),
		)
		if err != nil {
			return nil, err
		}

		a.syncProducer = p
	}

	return a.syncProducer, nil
}

func (a *App) GetOrderProducer(ctx context.Context) (kafka3.Producer, error) {
	if a.orderProducer == nil {
		syncProducer, err := a.GetSyncProducer(ctx)
		if err != nil {
			return nil, err
		}

		a.orderProducer = producer.NewProducer(
			syncProducer,
			config.AppConfig().OrderPaidProducerConfig.Topic(),
			logger.Logger(),
		)

		closer.AddNamed("kafka producer", func(ctx context.Context) error {
			return syncProducer.Close()

		})
	}

	return a.orderProducer, nil
}

func (a *App) initLogger(ctx context.Context) error {
	level := config.AppConfig().Logger.Level()
	asJson := config.AppConfig().Logger.AsJson()
	return logger.Init(
		level,
		asJson,
	)
}

func (a *App) initCloser(_ context.Context) error {
	closer.SetLogger(logger.Logger())
	return nil
}
