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
	"github.com/Muvi7z/boilerplate/notification/internal/service/consumer/order_paid_consumer"
	telegram2 "github.com/Muvi7z/boilerplate/notification/internal/service/telegram"
	kafka3 "github.com/Muvi7z/boilerplate/platform/kafka"
	"github.com/Muvi7z/boilerplate/platform/kafka/consumer"
	"github.com/Muvi7z/boilerplate/platform/logger"
	kafka2 "github.com/Muvi7z/boilerplate/platform/middleware/kafka"
	"github.com/go-telegram/bot"
)

type App struct {
	telegramClient  http.TelegramClient
	telegramService service.TelegramService
	telegramBot     *bot.Bot

	orderAssembledConsumerService service.OrderAssembledService
	orderPaidConsumerService      service.OrderPaidService
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

func (a *App) InitDeps(ctx context.Context) error {
	inits := []func(ctx context.Context) error{}

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

		a.orderPaidConsumerGroup = consumerGroup
	}

	return a.orderAssembledConsumerGroup, nil
}

func (a *App) getOrderAssembledDecoder(ctx context.Context) kafka.OrderAssemblyDecoder {
	if a.orderDecoder == nil {
		a.orderDecoder = decoder.NewOrderDecoder()
	}

	return a.orderDecoder
}
