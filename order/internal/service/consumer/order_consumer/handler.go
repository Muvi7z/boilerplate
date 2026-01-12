package order_consumer

import (
	"context"
	"github.com/Muvi7z/boilerplate/order/internal/entity"
	"github.com/Muvi7z/boilerplate/platform/kafka/consumer"
	"github.com/Muvi7z/boilerplate/platform/logger"
	"go.uber.org/zap"
)

func (s *service) OrderAssembledHandler(ctx context.Context, msg consumer.Message) error {
	event, err := s.orderDecoder.DecodeShipAssembled(msg.Value)
	if err != nil {
		logger.Error(ctx, "failed to decode ship assembled", zap.Error(err))
		return err
	}

	err = s.orderUseCase.UpdateStatusOrder(ctx, event.OrderUuid, entity.Completed)
	if err != nil {
		logger.Error(ctx, "failed to update status order", zap.Error(err))
		return err
	}

	return nil
}
