package repository

import (
	"context"
	"github.com/Muvi7z/boilerplate/inventory/internal/entity"
)

type PartRepository interface {
	GetPart(ctx context.Context, uuid string) (entity.Part, error)
	ListPart(ctx context.Context, filter entity.Filter) ([]entity.Part, error)
}
