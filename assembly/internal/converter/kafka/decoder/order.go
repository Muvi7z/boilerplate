package decoder

import (
	"fmt"
	"github.com/Muvi7z/boilerplate/assembly/internal/entity"
	events_v1 "github.com/Muvi7z/boilerplate/shared/pkg/proto/events/v1"
	"google.golang.org/protobuf/proto"
)

type decoder struct {
}

func NewOrderDecoder() *decoder {
	return &decoder{}
}

func (d *decoder) DecodeOrderPaid(data []byte) (entity.OrderPaidEvent, error) {
	var pb events_v1.OrderPaid
	if err := proto.Unmarshal(data, &pb); err != nil {
		return entity.OrderPaidEvent{}, fmt.Errorf("failed to unmarshal protobuf: %w", err)
	}

	return entity.OrderPaidEvent{
		EventUuid:       pb.EventUuid,
		OrderUuid:       pb.OrderUuid,
		UserUuid:        pb.UserUuid,
		PaymentMethod:   pb.PaymentMethod,
		TransactionUuid: pb.TransactionUuid,
	}, nil
}

func (d *decoder) DecodeShipAssembled(data []byte) (entity.ShipAssembledEvent, error) {
	var pb events_v1.ShipAssembled

	if err := proto.Unmarshal(data, &pb); err != nil {
		return entity.ShipAssembledEvent{}, fmt.Errorf("failed to unmarshal protobuf: %w", err)
	}

	return entity.ShipAssembledEvent{
		EventUuid:    pb.EventUuid,
		OrderUuid:    pb.OrderUuid,
		UserUuid:     pb.UserUuid,
		BuildTimeSec: pb.BuildTimeSec,
	}, nil
}
