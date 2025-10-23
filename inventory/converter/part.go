package converter

import (
	"github.com/Muvi7z/boilerplate/inventory/entity"
	inventory_v1 "github.com/Muvi7z/boilerplate/shared/pkg/proto/inventory/v1"
)

func InventoryPartInfoToPart(part *inventory_v1.PartInfo) entity.Part {
	//metadata := map[string]*entity.Value
	//if part.Metadata != nil {
	//	metadata = part.Metadata[''].GetOne()
	//}

	return entity.Part{
		Uuid:        part.Uuid,
		Name:        part.Name,
		Description: part.Description,
		Price:       part.Price,
		Category:    part.Category.String(),
		Dimensions: entity.Dimensions{
			Length: part.Dimensions.Length,
			Width:  part.Dimensions.Width,
			Height: part.Dimensions.Height,
			Weight: part.Dimensions.Weight,
		},
		Manufacturer: entity.Manufacturer{
			Name:    part.Manufacturer.Name,
			Country: part.Manufacturer.Country,
			Website: part.Manufacturer.Website,
		},
		Tags:      part.Tags,
		Metadata:  part.Metadata,
		CreatedAt: part.CreatedAt.AsTime(),
		UpdatedAt: part.UpdatedAt.AsTime(),
	}
}
