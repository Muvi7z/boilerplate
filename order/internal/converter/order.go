package converter

import (
	"github.com/Muvi7z/boilerplate/order/internal/entity"
	generated "github.com/Muvi7z/boilerplate/shared/pkg/server"
)

func GeneratedToCreateOrderEntity(gen generated.PostApiV1OrdersJSONBody) entity.CreateOrder {
	var parts []string

	if gen.PartUuids != nil {
		parts = *gen.PartUuids
	}

	return entity.CreateOrder{
		UserUuid:  gen.UserUuid.String(),
		PartUuids: parts,
	}
}

func OrderEntityFromGenerated(order entity.Order) generated.Order {
	var orderUuid generated.OrderUuid
	orderUuid = generated.OrderUuid([]byte(order.OrderUuid))

	var status generated.Status
	status = generated.Status(order.Status)

	var paymentMethod generated.PaymentMethod
	paymentMethod = generated.PaymentMethod(order.PaymentMethod)

	var transactionUuid generated.TransactionUuid
	transactionUuid = generated.TransactionUuid([]byte(order.TransactionUuid))

	var userUuid generated.UserUuid
	userUuid = generated.UserUuid([]byte(order.UserUuid))

	return generated.Order{
		OrderUuid:       &orderUuid,
		PartUuids:       &order.PartUuids,
		PaymentMethod:   &paymentMethod,
		Status:          &status,
		TotalPrice:      &order.TotalPrice,
		TransactionUuid: &transactionUuid,
		UserUuid:        &userUuid,
	}

}

func OrderEntityToGenerated(gen generated.Order) entity.Order {
	var paymentMethod string
	if gen.PaymentMethod != nil {
		paymentMethod = string(*gen.PaymentMethod)
	}
	var totalPrice float64
	if gen.TotalPrice != nil {
		totalPrice = *gen.TotalPrice
	}

	var status string
	if gen.Status != nil {
		status = string(*gen.Status)
	}

	var parts []string

	if gen.PartUuids != nil {
		parts = *gen.PartUuids
	}

	return entity.Order{
		OrderUuid:       gen.OrderUuid.String(),
		UserUuid:        gen.UserUuid.String(),
		PaymentMethod:   paymentMethod,
		PartUuids:       parts,
		TotalPrice:      totalPrice,
		TransactionUuid: gen.TransactionUuid.String(),
		Status:          status,
	}
}
