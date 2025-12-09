package converter

import (
	"github.com/Muvi7z/boilerplate/order/internal/entity"
	payment_v1 "github.com/Muvi7z/boilerplate/shared/pkg/proto/payment/v1"
	generated "github.com/Muvi7z/boilerplate/shared/pkg/server"
)

func GeneratedPayOrderToEntity(genOrder generated.PostApiV1OrdersOrderUuidPayJSONBody) entity.PayOrder {
	var payMethod string
	if genOrder.PaymentMethod != nil {
		payMethod = string(*genOrder.PaymentMethod)
	}

	return entity.PayOrder{
		OrderUuid:     "",
		UserUuid:      "",
		PaymentMethod: payMethod,
	}
}

func StringToPayOrderPaymentMethod(paymentMethod string) payment_v1.PaymentMethod {
	var res payment_v1.PaymentMethod

	switch paymentMethod {
	case entity.CARD:
		res = payment_v1.PaymentMethod_CARD
	case entity.SBP:
		res = payment_v1.PaymentMethod_SBP
	case entity.CreditCard:
		res = payment_v1.PaymentMethod_CREDIT_CARD
	case entity.InvestorMoney:
		res = payment_v1.PaymentMethod_INVESTOR_MONEY
	default:
		res = payment_v1.PaymentMethod_UNKNOWN
	}

	return res
}

func EntityPayOrderToPaymentPayOrder(payOrder entity.PayOrder) *payment_v1.PayOrderRequest {

	return &payment_v1.PayOrderRequest{
		OrderUuid:     payOrder.OrderUuid,
		UserUuid:      payOrder.UserUuid,
		PaymentMethod: StringToPayOrderPaymentMethod(payOrder.PaymentMethod),
	}
}
