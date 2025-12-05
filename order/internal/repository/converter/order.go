package converter

import (
	"github.com/Muvi7z/boilerplate/order/internal/entity"
	repoentity "github.com/Muvi7z/boilerplate/order/internal/repository/entity"
)

func RepositoryToOrderEntity(order repoentity.Order) entity.Order {
	//dst := make([]string, 0, len(order.PartUuids))
	//for _, s := range order.PartUuids {
	//	if s != nil {
	//		dst = append(dst, *s)
	//	}
	//}

	var transactionUuid string
	if order.TransactionUuid != nil {
		transactionUuid = *order.TransactionUuid
	}

	var paymentMethod string
	if order.PaymentMethod != nil {
		paymentMethod = *order.PaymentMethod
	}

	return entity.Order{
		OrderUuid:       order.OrderUuid,
		UserUuid:        order.UserUuid,
		PartUuids:       order.PartUuids,
		TotalPrice:      order.TotalPrice,
		TransactionUuid: transactionUuid,
		PaymentMethod:   paymentMethod,
		Status:          order.Status,
	}
}

func EntityToOrderRepository(order entity.Order) repoentity.Order {
	//dst := make([]*string, len(order.PartUuids))
	//for i := range order.PartUuids {
	//	s := order.PartUuids[i] // копируем значение
	//	dst[i] = &s
	//}
	return repoentity.Order{
		OrderUuid:       order.OrderUuid,
		UserUuid:        order.UserUuid,
		PartUuids:       order.PartUuids,
		TotalPrice:      order.TotalPrice,
		TransactionUuid: &order.TransactionUuid,
		PaymentMethod:   &order.PaymentMethod,
		Status:          order.Status,
	}
}
