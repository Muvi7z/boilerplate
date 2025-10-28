package converter

import (
	"github.com/Muvi7z/boilerplate/inventory/internal/entity"
	repoEntity "github.com/Muvi7z/boilerplate/inventory/internal/repository/entity"
)

func PartInfoToRepoEntity(info entity.Part) repoEntity.Part {
	return repoEntity.Part{
		Uuid:        info.Uuid,
		Name:        info.Name,
		Description: info.Description,
		Price:       info.Price,
		Category:    info.Category,
		Dimensions: repoEntity.DimensionsInfo{
			Length: info.Dimensions.Length,
			Width:  info.Dimensions.Width,
			Height: info.Dimensions.Height,
			Weight: info.Dimensions.Weight,
		},
		Manufacturer: repoEntity.ManufacturerInfo{
			Name:    info.Manufacturer.Name,
			Country: info.Manufacturer.Country,
			Website: info.Manufacturer.Website,
		},
		Tags:      info.Tags,
		Metadata:  info.Metadata,
		CreatedAt: info.CreatedAt,
		UpdatedAt: info.UpdatedAt,
	}
}

func RepoEntityToPartInfo(info repoEntity.Part) entity.Part {
	return entity.Part{
		Uuid:        info.Uuid,
		Name:        info.Name,
		Description: info.Description,
		Price:       info.Price,
		Category:    info.Category,
		Dimensions: entity.Dimensions{
			Length: info.Dimensions.Length,
			Width:  info.Dimensions.Width,
			Height: info.Dimensions.Height,
			Weight: info.Dimensions.Weight,
		},
		Manufacturer: entity.Manufacturer{
			Name:    info.Manufacturer.Name,
			Country: info.Manufacturer.Country,
			Website: info.Manufacturer.Website,
		},
		Tags:      info.Tags,
		Metadata:  info.Metadata,
		CreatedAt: info.CreatedAt,
		UpdatedAt: info.UpdatedAt,
	}
}

func ArrayPartInfoToRepoEntity(info []entity.Part) []repoEntity.Part {
	convertedParts := make([]repoEntity.Part, len(info))

	for i, _ := range info {
		convertedParts[i] = PartInfoToRepoEntity(info[i])
	}

	return convertedParts
}

func ArrayRepoEntityToPartInfo(info []repoEntity.Part) []entity.Part {
	convertedParts := make([]entity.Part, len(info))
	for i, _ := range info {
		convertedParts[i] = RepoEntityToPartInfo(info[i])
	}

	return convertedParts
}
