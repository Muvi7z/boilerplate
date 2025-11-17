package converter

import (
	"github.com/Muvi7z/boilerplate/order/internal/entity"
	repoentity "github.com/Muvi7z/boilerplate/order/internal/repository/entity"
)

func RepositoryToOrderEntity(order repoentity.Order) entity.Order {
	return entity.Order{
		OrderUuid:       order.OrderUuid,
		UserUuid:        order.UserUuid,
		PartUuids:       order.PartUuids,
		TotalPrice:      order.TotalPrice,
		TransactionUuid: order.TransactionUuid,
		PaymentMethod:   order.PaymentMethod,
		Status:          order.Status,
	}
}

func EntityToOrderRepository(order entity.Order) repoentity.Order {
	return repoentity.Order{
		OrderUuid:       order.OrderUuid,
		UserUuid:        order.UserUuid,
		PartUuids:       order.PartUuids,
		TotalPrice:      order.TotalPrice,
		TransactionUuid: order.TransactionUuid,
		PaymentMethod:   order.PaymentMethod,
		Status:          order.Status,
	}
}
