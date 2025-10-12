package order

import generated "github.com/Muvi7z/boilerplate/shared/pkg/server"

type UseCase struct {
	orders map[string]generated.GetApiV1OrdersOrderUuid200JSONResponse
}

func New() *UseCase {
	return &UseCase{}
}
