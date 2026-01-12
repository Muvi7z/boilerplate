package service

import (
	"context"
	"github.com/Muvi7z/boilerplate/notification/internal/entity"
)

type TelegramService interface {
	SendNotificationAssembly(ctx context.Context, event entity.ShipAssembledEvent) error
	SendNotificationPaid(ctx context.Context, event entity.OrderPaidEvent) error
}

type OrderAssembledService interface {
	RunConsumer(ctx context.Context) error
}

type OrderPaidService interface {
	RunConsumer(ctx context.Context) error
}
