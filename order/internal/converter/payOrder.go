package converter

import (
	"github.com/Muvi7z/boilerplate/order/internal/entity"
	generated "github.com/Muvi7z/boilerplate/shared/pkg/server"
)

func GeneratedPayOrderToEntity(genOrder generated.PostApiV1OrdersOrderUuidPayJSONBody) entity.PayOrder {
	payMethod := string(*genOrder.PaymentMethod)

	return entity.PayOrder{
		OrderUuid:     "",
		UserUuid:      "",
		PaymentMethod: payMethod,
	}
}
