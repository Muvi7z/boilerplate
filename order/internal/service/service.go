package service

import (
	"context"
	"github.com/Muvi7z/boilerplate/order/internal/entity"
)

type OrderConsumerService interface {
	RunConsumer(context.Context) error
}

type OrderProducerService interface {
	OrderPaid(ctx context.Context, event entity.OrderPaidEvent) error
}
