package converter

import (
	"github.com/Muvi7z/boilerplate/payment/internal/entity"
	payment_v1 "github.com/Muvi7z/boilerplate/shared/pkg/proto/payment/v1"
)

func PaymentToPaymentEntity(payment *payment_v1.PayOrderRequest) entity.Payment {

	return entity.Payment{
		OrderUUID:     payment.GetOrderUuid(),
		UserUUID:      payment.GetUserUuid(),
		PaymentMethod: payment.PaymentMethod.String(),
	}
}
