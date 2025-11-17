package repository

import (
	"context"
	"github.com/Muvi7z/boilerplate/order/internal/entity"
	"github.com/Muvi7z/boilerplate/order/internal/repository/converter"
	"github.com/google/uuid"
)

func (r *Repository) Create(ctx context.Context, order entity.Order) (string, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	id := uuid.New().String()

	order.OrderUuid = id

	r.orders[id] = converter.EntityToOrderRepository(order)

	return id, nil
}

func (r *Repository) Get(ctx context.Context, id string) (entity.Order, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	order, ok := r.orders[id]
	if !ok {
		return entity.Order{}, entity.ErrOrderNotFound
	}

	return converter.RepositoryToOrderEntity(order), nil
}

func (r *Repository) Update(ctx context.Context, id string, updateOrder entity.Order) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	repoOrder, ok := r.orders[id]
	if !ok {
		return entity.ErrOrderNotFound
	}

	if updateOrder.UserUuid != "" {
		repoOrder.UserUuid = updateOrder.UserUuid
	}

	if updateOrder.OrderUuid != "" {
		repoOrder.OrderUuid = updateOrder.OrderUuid
	}

	if updateOrder.TransactionUuid != "" {
		repoOrder.TransactionUuid = updateOrder.TransactionUuid
	}

	if updateOrder.PaymentMethod != "" {
		repoOrder.PaymentMethod = updateOrder.PaymentMethod
	}

	if updateOrder.PartUuids != nil {
		repoOrder.PartUuids = updateOrder.PartUuids
	}

	if updateOrder.Status != "" {
		repoOrder.Status = updateOrder.Status
	}

	if updateOrder.TotalPrice != 0 {
		repoOrder.TotalPrice = updateOrder.TotalPrice
	}

	r.orders[id] = repoOrder

	return nil
}
