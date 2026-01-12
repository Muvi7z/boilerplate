package order_producer

import (
	"context"
	"github.com/Muvi7z/boilerplate/order/internal/entity"
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
	return &service{orderPaidProducer}
}

func (s *service) OrderPaid(ctx context.Context, event entity.OrderPaidEvent) error {

	msg := &events_v1.OrderPaid{
		EventUuid:       event.EventUuid,
		OrderUuid:       event.OrderUuid,
		UserUuid:        event.UserUuid,
		PaymentMethod:   event.PaymentMethod,
		TransactionUuid: event.TransactionUuid,
	}

	payload, err := proto.Marshal(msg)
	if err != nil {
		logger.Error(ctx, "failed to marshal order paid", zap.Error(err))
		return err
	}

	err = s.orderPaidProducer.Send(ctx, []byte(event.EventUuid), payload)
	if err != nil {
		logger.Error(ctx, "failed to send order paid", zap.Error(err))
		return err
	}

	return nil
}
