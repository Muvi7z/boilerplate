package app

import (
	"context"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/Muvi7z/boilerplate/assembly/internal/config"
	kafka2 "github.com/Muvi7z/boilerplate/assembly/internal/converter/kafka"
	"github.com/Muvi7z/boilerplate/assembly/internal/converter/kafka/decoder"
	"github.com/Muvi7z/boilerplate/assembly/internal/service"
	"github.com/Muvi7z/boilerplate/assembly/internal/service/consumer/order_consumer"
	"github.com/Muvi7z/boilerplate/assembly/internal/service/producer/order_producer"
	"github.com/Muvi7z/boilerplate/platform/closer"
	kafka3 "github.com/Muvi7z/boilerplate/platform/kafka"
	"github.com/Muvi7z/boilerplate/platform/kafka/consumer"
	"github.com/Muvi7z/boilerplate/platform/kafka/producer"
	"github.com/Muvi7z/boilerplate/platform/logger"
	"github.com/Muvi7z/boilerplate/platform/middleware/kafka"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type App struct {
	orderConsumerService service.ConsumerService
	orderConsumer        kafka3.Consumer
	orderAssemblyDecoder kafka2.OrderAssemblyDecoder
	consumerGroup        sarama.ConsumerGroup

	orderProducerService service.OrderProducerService
	orderProducer        kafka3.Producer
	syncProducer         sarama.SyncProducer
}

func New(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.InitDeps(ctx)
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
		if err := a.runConsumer(ctx); err != nil {
			errCh <- errors.Errorf("consumer crashed: %v", err)
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

func (a *App) runConsumer(ctx context.Context) error {
	logger.Info(ctx, "Order Kafka consumer running")

	err := a.orderConsumerService.RunConsumer(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) InitDeps(ctx context.Context) error {
	inits := []func(ctx context.Context) error{
		a.initCloser,
		a.initLogger,
		a.initConsumerService,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) getConsumerGroup() (sarama.ConsumerGroup, error) {
	if a.consumerGroup == nil {
		consumerGroup, err := sarama.NewConsumerGroup(config.AppConfig().Kafka.Brokers(),
			config.AppConfig().OrderPaidConsumer.GroupId(),
			config.AppConfig().OrderPaidConsumer.Config())
		if err != nil {
			return nil, err
		}

		closer.AddNamed("Kafka consumer group", func(ctx context.Context) error {
			return a.consumerGroup.Close()
		})

		a.consumerGroup = consumerGroup
	}

	return a.consumerGroup, nil
}

func (a *App) initConsumerService(ctx context.Context) error {
	if a.orderConsumerService == nil {

		orderConsumer, err := a.getOrderConsumer(ctx)
		if err != nil {
			return err
		}

		orderProducerService := a.getOrderProducerService(ctx)

		a.orderConsumerService = order_consumer.NewService(
			orderConsumer,
			a.getOrderAssemblyDecoder(ctx),
			orderProducerService,
		)
	}

	return nil
}

func (a *App) getOrderConsumer(ctx context.Context) (kafka3.Consumer, error) {
	if a.orderConsumer == nil {

		consumerGroup, err := a.getConsumerGroup()
		if err != nil {
			return nil, err
		}

		orderConsumer := consumer.NewConsumer(
			consumerGroup,
			[]string{
				config.AppConfig().OrderPaidConsumer.Topic(),
			},
			logger.Logger(),
			kafka.Logging(logger.Logger()),
		)

		a.orderConsumer = orderConsumer
	}

	return a.orderConsumer, nil
}

func (a *App) getOrderProducerService(ctx context.Context) service.OrderProducerService {
	if a.orderProducerService == nil {
		orderProducer, err := a.getOrderProducer(ctx)
		if err != nil {
			panic(fmt.Sprintf("Error getting order producer service: %v", err))
		}

		orderProducerService := order_producer.NewService(orderProducer)

		a.orderProducerService = orderProducerService
	}

	return a.orderProducerService
}

func (a *App) getOrderProducer(ctx context.Context) (kafka3.Producer, error) {
	if a.orderProducer == nil {
		syncProducer := a.getSyncProducer()

		a.orderProducer = producer.NewProducer(
			syncProducer,
			config.AppConfig().OrderAssembledProducer.Topic(),
			logger.Logger(),
		)

		closer.AddNamed("Kafka producer", func(ctx context.Context) error {
			return a.syncProducer.Close()
		})
	}

	return a.orderProducer, nil
}

func (a *App) getSyncProducer() sarama.SyncProducer {
	if a.syncProducer == nil {
		p, err := sarama.NewSyncProducer(
			config.AppConfig().Kafka.Brokers(),
			config.AppConfig().OrderAssembledProducer.Config(),
		)
		if err != nil {
			panic(fmt.Sprintf("Failed to create sync producer: %s", err))
		}

		a.syncProducer = p
	}
	return a.syncProducer
}

func (a *App) getOrderAssemblyDecoder(ctx context.Context) kafka2.OrderAssemblyDecoder {
	if a.orderAssemblyDecoder == nil {
		a.orderAssemblyDecoder = decoder.NewOrderDecoder()
	}

	return a.orderAssemblyDecoder
}

func (a *App) initLogger(ctx context.Context) error {
	level := config.AppConfig().Logger.Level()
	asJson := config.AppConfig().Logger.AsJson()

	return logger.Init(level, asJson)
}

func (a *App) initCloser(ctx context.Context) error {
	closer.SetLogger(logger.Logger())

	return nil
}
