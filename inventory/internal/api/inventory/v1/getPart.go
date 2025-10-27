package v1

import (
	"context"
	"github.com/Muvi7z/boilerplate/inventory/converter"
	inventory_v1 "github.com/Muvi7z/boilerplate/shared/pkg/proto/inventory/v1"
)

func (a *api) GetPart(ctx context.Context, req *inventory_v1.GetPartRequest) (*inventory_v1.GetPartResponse, error) {
	partInfo, err := a.partService.GetPart(ctx, req.Uuid)
	if err != nil {
		return nil, err
	}

	partInventory, err := converter.PartToInventoryPartInfo(partInfo)
	if err != nil {
		return nil, err
	}

	return &inventory_v1.GetPartResponse{
		Part: partInventory,
	}, nil
}
