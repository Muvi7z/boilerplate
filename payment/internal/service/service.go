package service

import (
	"context"
	"github.com/Muvi7z/boilerplate/payment/internal/entity"
)

type Payment interface {
	PayOrder(ctx context.Context, payment entity.Payment) (string, error)
}
