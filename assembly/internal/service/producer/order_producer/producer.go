package order_producer

import (
	"context"
	"github.com/Muvi7z/boilerplate/assembly/internal/entity"
	"github.com/Muvi7z/boilerplate/platform/kafka"
	"github.com/Muvi7z/boilerplate/platform/logger"
	events_v1 "github.com/Muvi7z/boilerplate/shared/pkg/proto/events/v1"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

type service struct {
	orderPaidProducer kafka.Producer
}

func NewService(orderPaidProducer kafka.Producer) *service {
	return &service{orderPaidProducer: orderPaidProducer}
}

func (s *service) ProduceShipAssembled(ctx context.Context, event entity.ShipAssembledEvent) error {

	msg := &events_v1.ShipAssembled{
		EventUuid:    event.EventUuid,
		OrderUuid:    event.OrderUuid,
		UserUuid:     event.UserUuid,
		BuildTimeSec: event.BuildTimeSec,
	}

	payload, err := proto.Marshal(msg)
	if err != nil {
		logger.Error(ctx, "failed to marshal ShipAssembled", zap.Error(err))
		return err
	}

	err = s.orderPaidProducer.Send(ctx, []byte(event.EventUuid), payload)
	if err != nil {
		logger.Error(ctx, "failed to publish ship assembled event", zap.Error(err))
		return err
	}

	return nil
}
