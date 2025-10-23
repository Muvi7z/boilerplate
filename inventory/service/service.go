package service

import (
	"context"
	"github.com/Muvi7z/boilerplate/inventory/entity"
)

type PartService interface {
	GetPart(ctx context.Context, uuid string) (entity.Part, error)
}
