package grpc

import (
	"context"
	"github.com/Muvi7z/boilerplate/order/internal/entity"
	inventory_v1 "github.com/Muvi7z/boilerplate/shared/pkg/proto/inventory/v1"
)

type PaymentClient interface {
	PayOrder(ctx context.Context, payOrder entity.PayOrder) (string, error)
}

type InventoryClient interface {
	ListParts(ctx context.Context, filter entity.PartsFilter) (*inventory_v1.ListPartResponse, error)
}
