package v1

import (
	"context"
	inventory_v1 "github.com/Muvi7z/boilerplate/shared/pkg/proto/inventory/v1"
)

func (a *api) ListPart(ctx context.Context, listPart *inventory_v1.ListPartRequest) (*inventory_v1.ListPartResponse, error) {

	return &inventory_v1.ListPartResponse{}, nil
}
