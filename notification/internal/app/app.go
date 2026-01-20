package app

import (
	"context"
	"github.com/IBM/sarama"
	"github.com/Muvi7z/boilerplate/notification/internal/client/http"
	"github.com/Muvi7z/boilerplate/notification/internal/client/http/telegram"
	"github.com/Muvi7z/boilerplate/notification/internal/config"
	"github.com/Muvi7z/boilerplate/notification/internal/converter/kafka"
	"github.com/Muvi7z/boilerplate/notification/internal/converter/kafka/decoder"
	"github.com/Muvi7z/boilerplate/notification/internal/service"
	"github.com/Muvi7z/boilerplate/notification/internal/service/consumer/order_assembled_consumer"
	"github.com/Muvi7z/boilerplate/notification/internal/service/consumer/order_paid_consumer"
	telegram2 "github.com/Muvi7z/boilerplate/notification/internal/service/telegram"
	"github.com/Muvi7z/boilerplate/platform/closer"
	kafka3 "github.com/Muvi7z/boilerplate/platform/kafka"
	"github.com/Muvi7z/boilerplate/platform/kafka/consumer"
	"github.com/Muvi7z/boilerplate/platform/logger"
	kafka2 "github.com/Muvi7z/boilerplate/platform/middleware/kafka"
	"github.com/go-telegram/bot"
	"go.uber.org/zap"
)

type App struct {
	telegramClient  http.TelegramClient
	telegramService service.TelegramService
	telegramBot     *bot.Bot

	orderAssembledConsumerService service.OrderAssembledConsumerService
	orderPaidConsumerService      service.OrderPaidConsumerService
	orderAssembledConsumer        kafka3.Consumer
	orderPaidConsumer             kafka3.Consumer
	orderPaidConsumerGroup        sarama.ConsumerGroup
	orderAssembledConsumerGroup   sarama.ConsumerGroup
	orderDecoder                  kafka.OrderAssemblyDecoder
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
	errCh := make(chan error, 3)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go func() {
		if err := a.RunOrderAssemblerConsumer(ctx); err != nil {
			errCh <- err
		}
	}()

	go func() {
		if err := a.RunOrderPaidConsumer(ctx); err != nil {
			errCh <- err
		}
	}()

	select {
	case <-ctx.Done():
		logger.Info(ctx, "shutdown signal received")
	case err := <-errCh:
		logger.Error(ctx, "component crushed", zap.Error(err))

		cancel()

		<-ctx.Done()

		return err

	}

	return nil
}

func (a *App) RunOrderPaidConsumer(ctx context.Context) error {
	err := a.orderPaidConsumerService.RunConsumer(ctx)
	if err != nil {
		return err
	}

	return nil
}
func (a *App) RunOrderAssemblerConsumer(ctx context.Context) error {
	err := a.orderAssembledConsumerService.RunConsumer(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) InitDeps(ctx context.Context) error {
	inits := []func(ctx context.Context) error{
		a.initCloser,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initOrderPaidConsumerService(ctx context.Context) error {
	if a.orderPaidConsumerService == nil {
		orderDecoder := a.getOrderAssembledDecoder(ctx)

		orderConsumer, err := a.getOrderPaidConsumer(ctx)
		if err != nil {
			return err
		}

		telegramService, err := a.getTelegramService(ctx)
		if err != nil {
			return err
		}

		a.orderPaidConsumerService = order_paid_consumer.NewService(orderDecoder, orderConsumer, telegramService)
	}
	return nil
}

func (a *App) getTelegramBot(ctx context.Context) (*bot.Bot, error) {
	if a.telegramBot == nil {
		b, err := bot.New(
			config.AppConfig().TelegramConfig.Token(),
		)
		if err != nil {
			return nil, err
		}

		a.telegramBot = b
	}

	return a.telegramBot, nil
}

func (a *App) getTelegramService(ctx context.Context) (service.TelegramService, error) {
	if a.telegramService == nil {
		tgClient, err := a.getTelegramClient(ctx)
		if err != nil {
			return nil, err
		}

		a.telegramService = telegram2.NewService(tgClient)
	}

	return a.telegramService, nil
}

func (a *App) getTelegramClient(ctx context.Context) (http.TelegramClient, error) {
	if a.telegramClient == nil {
		botTelegram, err := a.getTelegramBot(ctx)
		if err != nil {
			return nil, err
		}

		a.telegramClient = telegram.NewClient(botTelegram)
	}

	return a.telegramClient, nil
}

func (a *App) getOrderPaidConsumer(ctx context.Context) (kafka3.Consumer, error) {
	if a.orderPaidConsumer == nil {
		consumerGroup, err := a.getOrderPaidConsumerGroup(ctx)
		if err != nil {
			return nil, err
		}

		orderConsumer := consumer.NewConsumer(
			consumerGroup,
			[]string{
				config.AppConfig().OrderPaidConsumerConfig.Topic(),
			},
			logger.Logger(),
			kafka2.Logging(logger.Logger()),
		)

		a.orderPaidConsumer = orderConsumer
	}

	return a.orderPaidConsumer, nil
}

func (a *App) getOrderAssembledConsumerGroup(ctx context.Context) (sarama.ConsumerGroup, error) {
	if a.orderAssembledConsumerGroup == nil {
		consumerGroup, err := sarama.NewConsumerGroup(
			config.AppConfig().KafkaConfig.Brokers(),
			config.AppConfig().OrderAssembledConsumerConfig.GroupId(),
			config.AppConfig().OrderAssembledConsumerConfig.Config(),
		)
		if err != nil {
			return nil, err
		}

		a.orderAssembledConsumerGroup = consumerGroup
	}

	return a.orderAssembledConsumerGroup, nil
}

func (a *App) getOrderAssembledConsumer(ctx context.Context) (kafka3.Consumer, error) {
	if a.orderAssembledConsumer == nil {
		consumerGroup, err := a.getOrderAssembledConsumerGroup(ctx)
		if err != nil {
			return nil, err
		}

		orderConsumer := consumer.NewConsumer(
			consumerGroup,
			[]string{
				config.AppConfig().OrderAssembledConsumerConfig.Topic(),
			},
			logger.Logger(),
		)

		a.orderAssembledConsumer = orderConsumer
	}

	return a.orderAssembledConsumer, nil
}

func (a *App) initOrderAssembledConsumerService(ctx context.Context) (service.OrderAssembledConsumerService, error) {
	if a.orderAssembledConsumerService == nil {
		orderConsumer, err := a.getOrderAssembledConsumer(ctx)
		if err != nil {
			return nil, err
		}

		tgService, err := a.getTelegramService(ctx)
		if err != nil {
			return nil, err
		}

		a.orderAssembledConsumerService = order_assembled_consumer.NewService(orderConsumer, a.getOrderAssembledDecoder(ctx), tgService)
	}

	return a.orderAssembledConsumerService, nil
}

func (a *App) getOrderPaidConsumerGroup(ctx context.Context) (sarama.ConsumerGroup, error) {
	if a.orderPaidConsumerGroup == nil {
		consumerGroup, err := sarama.NewConsumerGroup(
			config.AppConfig().KafkaConfig.Brokers(),
			config.AppConfig().OrderPaidConsumerConfig.GroupId(),
			config.AppConfig().OrderPaidConsumerConfig.Config(),
		)
		if err != nil {
			return nil, err
		}

		a.orderPaidConsumerGroup = consumerGroup
	}

	return a.orderPaidConsumerGroup, nil
}

func (a *App) getOrderAssembledDecoder(ctx context.Context) kafka.OrderAssemblyDecoder {
	if a.orderDecoder == nil {
		a.orderDecoder = decoder.NewOrderDecoder()
	}

	return a.orderDecoder
}

func (a *App) initCloser(_ context.Context) error {
	closer.SetLogger(logger.Logger())
	return nil
}
