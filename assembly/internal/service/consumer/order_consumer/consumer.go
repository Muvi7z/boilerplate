package order_consumer

import (
	"context"
	kafka2 "github.com/Muvi7z/boilerplate/assembly/internal/converter/kafka"
	service2 "github.com/Muvi7z/boilerplate/assembly/internal/service"
	"github.com/Muvi7z/boilerplate/platform/kafka"
	"github.com/Muvi7z/boilerplate/platform/logger"
)

type service struct {
	orderConsumer                 kafka.Consumer
	orderAssemblyDecoder          kafka2.OrderAssemblyDecoder
	orderAssembledProducerService service2.OrderProducerService
}

func NewService(orderConsumer kafka.Consumer, orderAssemblyDecoder kafka2.OrderAssemblyDecoder, orderProducerService service2.OrderProducerService) *service {
	return &service{
		orderConsumer:                 orderConsumer,
		orderAssemblyDecoder:          orderAssemblyDecoder,
		orderAssembledProducerService: orderProducerService,
	}
}

func (s *service) RunConsumer(ctx context.Context) error {
	logger.Info(ctx, "Starting order consumer service")

	err := s.orderConsumer.Consume(ctx, s.OrderHandler)
	if err != nil {
		logger.Error(ctx, "consume from order.paid topic error")
		return err
	}

	return nil
}
