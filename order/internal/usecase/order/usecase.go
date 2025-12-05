package order

import (
	"context"
	"github.com/Muvi7z/boilerplate/order/internal/client/converter"
	"github.com/Muvi7z/boilerplate/order/internal/client/grpc"
	"github.com/Muvi7z/boilerplate/order/internal/entity"
	generated "github.com/Muvi7z/boilerplate/shared/pkg/server"
)

type UseCase struct {
	paymentClient   grpc.PaymentClient
	inventoryClient grpc.InventoryClient
	orderRepository Repository
}

func New(paymentClient grpc.PaymentClient, inventoryClient grpc.InventoryClient, orderRepository Repository) *UseCase {
	return &UseCase{
		paymentClient:   paymentClient,
		orderRepository: orderRepository,
		inventoryClient: inventoryClient,
	}
}

func (u *UseCase) GetOrder(ctx context.Context, id string) (entity.Order, error) {
	order, err := u.orderRepository.Get(ctx, id)
	if err != nil {
		return entity.Order{}, err
	}

	return *order, nil
}

func (u *UseCase) CreateOrder(ctx context.Context, createOrder entity.CreateOrder) (*entity.Order, error) {

	filter := entity.PartsFilter{
		Uuids:                 createOrder.PartUuids,
		Names:                 nil,
		Categories:            nil,
		ManufacturerCountries: nil,
		Tags:                  nil,
	}

	inventoryParts, err := u.inventoryClient.ListParts(ctx, filter)
	if err != nil {
		return nil, err
	}

	parts := converter.InventoryPartInfoToArrayPartInfoEntity(inventoryParts.Parts)

	var totalPrice float64

	for _, part := range parts {
		totalPrice = totalPrice + part.Price
	}

	order := entity.Order{
		UserUuid:        createOrder.UserUuid,
		PartUuids:       createOrder.PartUuids,
		TotalPrice:      totalPrice,
		TransactionUuid: "",
		PaymentMethod:   "",
		Status:          string(generated.PENDINGPAYMENT),
	}

	uuidOrder, err := u.orderRepository.Create(ctx, order)
	if err != nil {
		return nil, err
	}

	order.OrderUuid = uuidOrder

	return &order, nil
}

func (u *UseCase) CancelOrder(ctx context.Context, id string) error {
	order, err := u.orderRepository.Get(ctx, id)
	if err != nil {
		return err
	}

	if order.Status == string(generated.PAID) {
		return entity.ErrOrderIsPaid
	}

	order.Status = string(generated.PENDINGPAYMENT)

	err = u.orderRepository.Update(ctx, order.OrderUuid, *order)
	if err != nil {
		return err
	}

	return nil
}

func (u *UseCase) PayOrder(ctx context.Context, payOrder entity.PayOrder) (string, error) {

	transactionUuid, err := u.paymentClient.PayOrder(ctx, payOrder)
	if err != nil {
		return "", err
	}

	order := entity.Order{
		OrderUuid:       payOrder.OrderUuid,
		UserUuid:        payOrder.UserUuid,
		PaymentMethod:   payOrder.PaymentMethod,
		TransactionUuid: transactionUuid,
		Status:          string(generated.PAID),
	}

	err = u.orderRepository.Update(ctx, payOrder.OrderUuid, order)
	if err != nil {
		return "", err
	}

	return transactionUuid, nil
}
