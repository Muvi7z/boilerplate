package order_consumer

import (
	"context"
	"github.com/Muvi7z/boilerplate/assembly/internal/entity"
	"github.com/Muvi7z/boilerplate/platform/kafka/consumer"
	"github.com/Muvi7z/boilerplate/platform/logger"
	"go.uber.org/zap"
	"math/rand"
	"time"
)

func (s *service) OrderHandler(ctx context.Context, msg consumer.Message) error {
	event, err := s.orderAssemblyDecoder.DecodeOrderPaid(msg.Value)
	if err != nil {
		logger.Error(ctx, "failed to decode order paid", zap.Error(err))
		return err
	}

	d := rand.Int()%10 + 1

	time.Sleep(time.Duration(d) * time.Second)

	err = s.orderAssembledProducerService.ProduceShipAssembled(ctx, entity.ShipAssembledEvent{
		EventUuid:    event.EventUuid,
		OrderUuid:    event.OrderUuid,
		UserUuid:     event.UserUuid,
		BuildTimeSec: int64(d),
	})
	if err != nil {
		logger.Error(ctx, "failed to send ship assembled", zap.Error(err))
		return err
	}

	return nil
}
