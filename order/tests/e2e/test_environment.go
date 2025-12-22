package integration

import (
	"context"
	"fmt"
	"github.com/Muvi7z/boilerplate/order/internal/entity"
	generated "github.com/Muvi7z/boilerplate/shared/pkg/server"
	"github.com/brianvoe/gofakeit/v7"
)

//
//func (env *TestEnvironment) InsertTestSighting(ctx context.Context)(string, error)  {
//	orderUUID := gofakeit.UUID()
//	now := time.Now()
//
//
//}

func (env *TestEnvironment) GetTestOrder() entity.Order {
	return entity.Order{
		OrderUuid:       gofakeit.UUID(),
		UserUuid:        gofakeit.UUID(),
		PartUuids:       []string{gofakeit.UUID(), gofakeit.UUID()},
		TotalPrice:      gofakeit.Price(10, 3500),
		TransactionUuid: "",
		PaymentMethod:   "",
		Status:          string(generated.PENDINGPAYMENT),
	}
}

func (env *TestEnvironment) ClearOrderTable(ctx context.Context) error {
	tableName := "orders"

	sql := fmt.Sprintf("TRUNCATE TABLE %s CASCADE", tableName)

	_, err := env.Postgres.Client().Exec(sql)
	if err != nil {
		return err
	}

	return nil
}
