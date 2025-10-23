package order

import (
	"context"
	generated "github.com/Muvi7z/boilerplate/shared/pkg/server"
	"github.com/google/uuid"
)

type UseCase struct {
	orders map[string]generated.GetApiV1OrdersOrderUuid200JSONResponse
}

func New() *UseCase {
	return &UseCase{}
}

func (u *UseCase) Create(ctx context.Context) error {
	id := uuid.New()

	//u.orders[id]

	return nil
}
