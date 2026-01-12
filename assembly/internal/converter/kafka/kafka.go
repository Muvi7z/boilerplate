package kafka

import "github.com/Muvi7z/boilerplate/assembly/internal/entity"

type OrderAssemblyDecoder interface {
	DecodeOrderPaid(data []byte) (entity.OrderPaidEvent, error)
	DecodeShipAssembled(data []byte) (entity.ShipAssembledEvent, error)
}
