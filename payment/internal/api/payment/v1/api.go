package v1

import (
	"github.com/Muvi7z/boilerplate/payment/internal/service"
	payment_v1 "github.com/Muvi7z/boilerplate/shared/pkg/proto/payment/v1"
)

type api struct {
	payment_v1.UnimplementedPaymentServer

	paymentService service.Payment
}

func NewAPI(paymentService service.Payment) *api {
	return &api{paymentService: paymentService}
}
