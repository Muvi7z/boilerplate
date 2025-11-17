package v1

import (
	"context"
	"github.com/Muvi7z/boilerplate/order/internal/client/converter"
	"github.com/Muvi7z/boilerplate/order/internal/entity"
	inventory_v1 "github.com/Muvi7z/boilerplate/shared/pkg/proto/inventory/v1"
)

func (c *client) ListParts(ctx context.Context, filter entity.PartsFilter) (*inventory_v1.ListPartResponse, error) {
	parts, err := c.grpcInventoryClient.ListPart(ctx, converter.PartsFilterToInventoryListPartRequest(filter))
	if err != nil {
		return &inventory_v1.ListPartResponse{}, err
	}

	return parts, nil
}
