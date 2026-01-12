package order_assembled_consumer

import (
	"context"
	"github.com/Muvi7z/boilerplate/notification/internal/converter/kafka"
	service2 "github.com/Muvi7z/boilerplate/notification/internal/service"
	kafka3 "github.com/Muvi7z/boilerplate/platform/kafka"
	"github.com/Muvi7z/boilerplate/platform/logger"
	"go.uber.org/zap"
)

type service struct {
	orderAssemblerConsumer kafka3.Consumer
	orderAssemblyDecoder   kafka.OrderAssemblyDecoder
	telegramService        service2.TelegramService
}

func NewService(orderAssemblerConsumer kafka3.Consumer, orderAssemblyDecoder kafka.OrderAssemblyDecoder, telegramService service2.TelegramService) *service {
	return &service{
		orderAssemblerConsumer: orderAssemblerConsumer,
		orderAssemblyDecoder:   orderAssemblyDecoder,
		telegramService:        telegramService,
	}
}

func (s *service) RunConsumer(ctx context.Context) error {
	logger.Info(ctx, "Starting order orderAssemblerConsumer service")

	err := s.orderAssemblerConsumer.Consume(ctx, s.OrderAssembledConsumer)
	if err != nil {
		logger.Error(ctx, "Consume from order.assembled topic error", zap.Error(err))

		return err
	}

	return nil
}
