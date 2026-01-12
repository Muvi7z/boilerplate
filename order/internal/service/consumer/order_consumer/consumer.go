package order_consumer

import (
	"context"
	kafka2 "github.com/Muvi7z/boilerplate/order/internal/converter/kafka"
	"github.com/Muvi7z/boilerplate/order/internal/usecase/order"
	"github.com/Muvi7z/boilerplate/platform/kafka"
	"github.com/Muvi7z/boilerplate/platform/logger"
)

type service struct {
	orderConsumer kafka.Consumer
	orderDecoder  kafka2.OrderAssemblyDecoder
	orderUseCase  *order.UseCase
}

func NewService(orderConsumer kafka.Consumer, orderDecoder kafka2.OrderAssemblyDecoder, orderUseCase *order.UseCase) *service {
	return &service{
		orderConsumer: orderConsumer,
		orderDecoder:  orderDecoder,
		orderUseCase:  orderUseCase,
	}
}

func (s *service) RunConsumer(ctx context.Context) error {
	logger.Info(ctx, "Starting order consumer")

	err := s.orderConsumer.Consume(ctx, s.OrderAssembledHandler)
	if err != nil {
		logger.Error(ctx, "consume from order.assembled")
		return err
	}

	return nil
}
