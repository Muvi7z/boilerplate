package payment

import (
	"context"
	"github.com/Muvi7z/boilerplate/payment/internal/entity"
	"github.com/google/uuid"
)

func (s *service) PayOrder(ctx context.Context, payment entity.Payment) (string, error) {
	id := uuid.New()

	return id.String(), nil
}
