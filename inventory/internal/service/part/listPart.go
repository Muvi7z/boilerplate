package part

import (
	"context"
	"github.com/Muvi7z/boilerplate/inventory/internal/entity"
)

func (s *service) ListPart(ctx context.Context, filter entity.Filter) ([]entity.Part, error) {
	listPart, err := s.partRepository.ListPart(ctx, filter)
	if err != nil {
		return nil, err
	}

	return listPart, nil
}
