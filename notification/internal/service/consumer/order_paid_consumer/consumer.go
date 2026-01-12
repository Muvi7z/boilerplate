package order_paid_consumer

import (
	"context"
	"github.com/Muvi7z/boilerplate/notification/internal/converter/kafka"
	service2 "github.com/Muvi7z/boilerplate/notification/internal/service"
	kafka2 "github.com/Muvi7z/boilerplate/platform/kafka"
	"github.com/Muvi7z/boilerplate/platform/logger"
	"go.uber.org/zap"
)

type service struct {
	orderAssemblyDecoder kafka.OrderAssemblyDecoder
	orderConsumer        kafka2.Consumer
	telegramService      service2.TelegramService
}

func NewService(orderAssemblyDecoder kafka.OrderAssemblyDecoder, orderConsumer kafka2.Consumer, telegramService service2.TelegramService) *service {
	return &service{
		orderAssemblyDecoder: orderAssemblyDecoder,
		orderConsumer:        orderConsumer,
		telegramService:      telegramService,
	}
}

func (s *service) RunConsumer(ctx context.Context) error {
	logger.Info(ctx, "Starting consumer orderPaid service")

	err := s.orderConsumer.Consume(ctx, s.OrderPaidConsumer)
	if err != nil {
		logger.Error(ctx, "Failed to consume from Kafka", zap.Error(err))
		return err
	}

	return nil
}
