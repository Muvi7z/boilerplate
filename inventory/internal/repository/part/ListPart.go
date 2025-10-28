package part

import (
	"context"
	"github.com/Muvi7z/boilerplate/inventory/internal/entity"
	"github.com/Muvi7z/boilerplate/inventory/internal/repository/converter"
	repoEntity "github.com/Muvi7z/boilerplate/inventory/internal/repository/entity"
	"strings"
)

func (r *repository) ListPart(ctx context.Context, filter entity.Filter) ([]entity.Part, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	filteredParts := make([]repoEntity.Part, len(r.parts), 0)

	for i, _ := range r.parts {
		accepted := false
		index := 0

		for _, fUuid := range filter.Uuids {
			if strings.Contains(fUuid, r.parts[i].Uuid) {
				accepted = true
			}
		}

		for _, fName := range filter.Names {
			if strings.Contains(fName, r.parts[i].Name) {
				accepted = true
			}
		}

		for _, fCategory := range filter.Categories {
			if strings.Contains(fCategory, r.parts[i].Category) {
				accepted = true
			}
		}

		for _, fManufacturerCountry := range filter.ManufacturerCountries {
			if strings.Contains(fManufacturerCountry, r.parts[i].Manufacturer.Country) {
				accepted = true
			}
		}

		for _, fTag := range filter.Tags {
			for _, tag := range r.parts[i].Tags {
				if strings.Contains(fTag, tag) {
					accepted = true
				}
			}

		}

		if len(filter.Uuids) == 0 && len(filter.Categories) == 0 && len(filter.ManufacturerCountries) == 0 &&
			len(filter.Names) == 0 && len(filter.Tags) == 0 {
			accepted = true
		}

		if accepted {
			filteredParts[index] = r.parts[i]
		}
	}

	return converter.ArrayRepoEntityToPartInfo(filteredParts), nil

}
