package order_assembled_consumer

import (
	"context"
	"github.com/Muvi7z/boilerplate/platform/kafka/consumer"
	"github.com/Muvi7z/boilerplate/platform/logger"
	"go.uber.org/zap"
)

func (s *service) OrderAssembledConsumer(ctx context.Context, msg consumer.Message) error {
	event, err := s.orderAssemblyDecoder.DecodeShipAssembled(msg.Value)
	if err != nil {
		logger.Error(ctx, "failed to decode order assembled", zap.Error(err))
		return err
	}

	err = s.telegramService.SendNotificationAssembly(ctx, event)
	if err != nil {
		logger.Error(ctx, "failed to send notification assembled message", zap.Error(err))
		return err
	}

	return nil
}
