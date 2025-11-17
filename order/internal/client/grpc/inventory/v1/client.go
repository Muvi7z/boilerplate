package v1

import (
	inventory_v1 "github.com/Muvi7z/boilerplate/shared/pkg/proto/inventory/v1"
)

type client struct {
	grpcInventoryClient inventory_v1.InventoryServiceClient
}

func New(grpcInventoryClient inventory_v1.InventoryServiceClient) *client {
	return &client{
		grpcInventoryClient: grpcInventoryClient,
	}
}
