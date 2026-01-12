package service

import (
	"context"
	"github.com/Muvi7z/boilerplate/assembly/internal/entity"
)

type ConsumerService interface {
	RunConsumer(ctx context.Context) error
}

type OrderProducerService interface {
	ProduceShipAssembled(ctx context.Context, event entity.ShipAssembledEvent) error
}
