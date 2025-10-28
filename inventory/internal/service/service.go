package service

import (
	"context"
	"github.com/Muvi7z/boilerplate/inventory/internal/entity"
)

type PartService interface {
	GetPart(ctx context.Context, uuid string) (entity.Part, error)
	ListPart(ctx context.Context, filter entity.Filter) ([]entity.Part, error)
}
