package app

import (
	"context"
	"github.com/IBM/sarama"
	"github.com/Muvi7z/boilerplate/notification/internal/client/http"
	"github.com/Muvi7z/boilerplate/notification/internal/converter/kafka"
	"github.com/Muvi7z/boilerplate/notification/internal/service"
	kafka3 "github.com/Muvi7z/boilerplate/platform/kafka"
)

type App struct {
	telegramClient  http.TelegramClient
	telegramService service.TelegramService

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
		NewService
	}
	return nil
}
