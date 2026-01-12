package order_paid_consumer

import (
	"context"
	"github.com/Muvi7z/boilerplate/platform/kafka/consumer"
	"github.com/Muvi7z/boilerplate/platform/logger"
	"go.uber.org/zap"
)

func (s *service) OrderPaidConsumer(ctx context.Context, msg consumer.Message) error {
	event, err := s.orderAssemblyDecoder.DecodeOrderPaid(msg.Value)
	if err != nil {
		logger.Info(ctx, "Failed to decode order paid event")
		return err
	}

	err = s.telegramService.SendNotificationPaid(ctx, event)
	if err != nil {
		logger.Error(ctx, "Failed to send notification paid event", zap.Error(err))
		return err
	}

	return nil
}
