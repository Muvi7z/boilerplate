package repository

import (
	"context"
	"errors"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/Muvi7z/boilerplate/order/internal/entity"
	"github.com/Muvi7z/boilerplate/order/internal/repository/converter"
	entity2 "github.com/Muvi7z/boilerplate/order/internal/repository/entity"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

const ordersTable = "orders"

func (r *Repository) Create(ctx context.Context, order entity.Order) (string, error) {
	var res string
	var err, txErr error

	txErr = sqlxTransaction(ctx, r.db, func(tx *sqlx.Tx) error {
		id := uuid.New().String()

		order.OrderUuid = id

		res, err = r.createTX(ctx, order, tx)

		return err
	})

	if txErr != nil {
		return "", txErr
	}

	return res, nil
}

func (r *Repository) createTX(ctx context.Context, order entity.Order, tx *sqlx.Tx) (string, error) {
	insertMap := map[string]any{
		"order_uuid":  order.OrderUuid,
		"part_uuids":  pq.Array(order.PartUuids),
		"user_uuid":   order.UserUuid,
		"total_price": order.TotalPrice,
	}

	if order.PaymentMethod != "" {
		insertMap["payment_method"] = order.PaymentMethod
	}

	if order.TransactionUuid != "" {
		insertMap["transaction_uuid"] = order.TransactionUuid
	}

	if order.Status != "" {
		insertMap["status"] = order.Status
	}

	sql, args, err := r.qb.Insert(ordersTable).
		SetMap(insertMap).
		Suffix("RETURNING *").
		ToSql()
	if err != nil {
		return "", fmt.Errorf("error building query: %w", err)
	}

	var row entity2.Order

	err = tx.GetContext(ctx, &row, sql, args...)
	if err != nil {
		return "", errors.Join(entity.ErrCreateOrder, err)
	}

	return row.OrderUuid, nil
}

func (r *Repository) Get(ctx context.Context, id string) (*entity.Order, error) {
	whereMap := map[string]any{
		"order_uuid": id,
	}

	sql, args, err := r.qb.Select("order_uuid").
		Columns("order_uuid", "user_uuid", "part_uuids", "total_price", "transaction_uuid", "payment_method", "status").
		From(ordersTable).
		Where(whereMap).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("error building query: %w", err)
	}

	var row entity2.Order
	err = r.db.GetContext(ctx, &row, sql, args...)
	if err != nil {
		return nil, errors.Join(entity.ErrGetOrder, err)
	}

	order := converter.RepositoryToOrderEntity(row)

	return &order, nil
}

func (r *Repository) Update(ctx context.Context, id string, updateOrder entity.Order) error {
	updateMap := map[string]any{}

	if updateOrder.TransactionUuid != "" {
		updateMap["transaction_uuid"] = updateOrder.TransactionUuid
	}

	if updateOrder.PaymentMethod != "" {
		updateMap["payment_method"] = updateOrder.PaymentMethod
	}

	if updateOrder.PartUuids != nil {
		updateMap["part_uuids"] = pq.Array(updateOrder.PartUuids)
	}

	if updateOrder.Status != "" {
		updateMap["status"] = updateOrder.Status
	}

	if updateOrder.TotalPrice != 0 {
		updateMap["total_price"] = updateOrder.TotalPrice
	}

	sql, ergs, err := r.qb.
		Update(ordersTable).
		SetMap(updateMap).
		Where(sq.Eq{"order_uuid": id}).
		Suffix("RETURNING *").
		ToSql()
	if err != nil {
		return fmt.Errorf("error building query: %w", err)
	}

	var row entity2.Order
	err = r.db.GetContext(ctx, &row, sql, ergs...)
	if err != nil {
		return errors.Join(entity.ErrUpdateOrder, err)
	}

	return nil
}
