package order

import (
	"context"
	"github.com/Muvi7z/boilerplate/order/internal/entity"
)

type Repository interface {
	Create(ctx context.Context, order entity.Order) (string, error)
	Update(ctx context.Context, id string, updateOrder entity.Order) error
	Get(ctx context.Context, id string) (entity.Order, error)
}
