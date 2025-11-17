package converter

import (
	"github.com/Muvi7z/boilerplate/order/internal/entity"
	inventory_v1 "github.com/Muvi7z/boilerplate/shared/pkg/proto/inventory/v1"
	"time"
)

func StringToCategory(category string) inventory_v1.Category {
	var res inventory_v1.Category

	switch category {
	case entity.CATEGORY_ENGINE:
		res = inventory_v1.Category_ENGINE
	case entity.CATEGORY_FUEL:
		res = inventory_v1.Category_FUEL
	case entity.CATEGORY_PORTHOLE:
		res = inventory_v1.Category_PORTHOLE
	case entity.CATEGORY_WING:
		res = inventory_v1.Category_WING
	default:
		res = inventory_v1.Category_UNKNOWN
	}

	return res
}

func PartsFilterToInventoryListPartRequest(filter entity.PartsFilter) *inventory_v1.ListPartRequest {
	var categories []inventory_v1.Category

	for _, category := range filter.Categories {
		categories = append(categories, StringToCategory(category))
	}

	return &inventory_v1.ListPartRequest{
		Filter: &inventory_v1.PartFilter{
			Uuids:                 filter.Uuids,
			Names:                 filter.Names,
			Categories:            categories,
			ManufacturerCountries: nil,
			Tags:                  nil,
		},
	}
}

func InventoryToPartInfoEntity(part *inventory_v1.PartInfo) entity.Part {
	//metadata := map[string]*entity.Value
	//if part.Metadata != nil {
	//	metadata = part.Metadata[''].GetOne()
	//}
	var createdAt time.Time
	var updatedAt time.Time
	if part.CreatedAt != nil {
		createdAt = part.CreatedAt.AsTime()
	}

	if part.UpdatedAt != nil {
		updatedAt = part.UpdatedAt.AsTime()
	}

	metadata := make(map[string]interface{})

	for key, val := range part.Metadata {
		switch v := val.One.(type) {
		case *inventory_v1.Value_StringValue:
			metadata[key] = v.StringValue
		case *inventory_v1.Value_Int64Value:
			metadata[key] = v.Int64Value
		case *inventory_v1.Value_DoubleValue:
			metadata[key] = v.DoubleValue
		case *inventory_v1.Value_BoolValue:
			metadata[key] = v.BoolValue
		default:
			metadata[key] = nil
		}
	}

	return entity.Part{
		Uuid:        part.Uuid,
		Name:        part.Name,
		Description: part.Description,
		Price:       part.Price,
		Category:    part.Category.String(),
		Dimensions: entity.DimensionsInfo{
			Length: part.Dimensions.Length,
			Width:  part.Dimensions.Width,
			Height: part.Dimensions.Height,
			Weight: part.Dimensions.Weight,
		},
		Manufacturer: entity.ManufacturerInfo{
			Name:    part.Manufacturer.Name,
			Country: part.Manufacturer.Country,
			Website: part.Manufacturer.Website,
		},
		Tags:      part.Tags,
		Metadata:  metadata,
		CreatedAt: &createdAt,
		UpdatedAt: &updatedAt,
	}
}

func InventoryPartInfoToArrayPartInfoEntity(info []*inventory_v1.PartInfo) []entity.Part {
	convertedParts := make([]entity.Part, len(info))
	for i, _ := range info {
		convertedParts[i] = InventoryToPartInfoEntity(info[i])
	}

	return convertedParts
}
