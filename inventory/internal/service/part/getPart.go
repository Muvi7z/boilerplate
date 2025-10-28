package part

import (
	"context"
	"github.com/Muvi7z/boilerplate/inventory/internal/entity"
)

func (s *service) GetPart(ctx context.Context, uuid string) (entity.Part, error) {
	part, err := s.partRepository.GetPart(ctx, uuid)
	if err != nil {
		return entity.Part{}, err
	}

	return part, nil
}
