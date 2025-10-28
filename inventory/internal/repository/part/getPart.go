package part

import (
	"context"
	"github.com/Muvi7z/boilerplate/inventory/internal/entity"
	"github.com/Muvi7z/boilerplate/inventory/internal/repository/converter"
)

func (r *repository) GetPart(ctx context.Context, uuid string) (entity.Part, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	part, ok := r.parts[uuid]
	if !ok {
		return entity.Part{}, entity.ErrPartInfoNotFound
	}

	return converter.RepoEntityToPartInfo(part), nil
}
