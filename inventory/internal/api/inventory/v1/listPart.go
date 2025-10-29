package v1

import (
	"context"
	"github.com/Muvi7z/boilerplate/inventory/internal/converter"
	inventory_v1 "github.com/Muvi7z/boilerplate/shared/pkg/proto/inventory/v1"
)

func (a *api) ListPart(ctx context.Context, listPart *inventory_v1.ListPartRequest) (*inventory_v1.ListPartResponse, error) {
	parts, err := a.partService.ListPart(ctx, converter.InventoryPartFilterToPartFilter(listPart.Filter))
	if err != nil {
		return nil, err
	}
	return &inventory_v1.ListPartResponse{
		Parts: converter.ArrayPartEntityToPartInfo(parts),
	}, nil
}
