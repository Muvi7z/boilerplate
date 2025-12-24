package app

import (
	"github.com/Muvi7z/boilerplate/order/internal/usecase/order"
)

type container struct {
	orderRepository order.Repository
}

func NewContainer() *container {
	return &container{}
}
