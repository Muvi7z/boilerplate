package converter

import (
	"github.com/Muvi7z/boilerplate/inventory/internal/entity"
	inventory_v1 "github.com/Muvi7z/boilerplate/shared/pkg/proto/inventory/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func InventoryPartInfoToPart(part *inventory_v1.PartInfo) entity.Part {
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
		Metadata:  metadata,
		CreatedAt: &createdAt,
		UpdatedAt: &updatedAt,
	}
}

func StringToCategory(category string) inventory_v1.Category {
	var res inventory_v1.Category

	switch category {
	case entity.ENGINE:
		res = inventory_v1.Category_ENGINE
	case entity.FUEL:
		res = inventory_v1.Category_FUEL
	case entity.PORTHOLE:
		res = inventory_v1.Category_PORTHOLE
	case entity.WING:
		res = inventory_v1.Category_WING
	default:
		res = inventory_v1.Category_UNKNOWN
	}

	return res
}

func PartToInventoryPartInfo(part entity.Part) *inventory_v1.PartInfo {
	var createdAt *timestamppb.Timestamp
	if part.CreatedAt != nil {
		createdAt = timestamppb.New(*part.CreatedAt)
	}
	var updatedAt *timestamppb.Timestamp
	if part.UpdatedAt != nil {
		updatedAt = timestamppb.New(*part.UpdatedAt)
	}

	metadata := make(map[string]*inventory_v1.Value)

	for key, val := range part.Metadata {
		switch v := val.(type) {
		case string:
			metadata[key] = &inventory_v1.Value{
				One: &inventory_v1.Value_StringValue{
					StringValue: v,
				},
			}
		case int64:
			metadata[key] = &inventory_v1.Value{
				One: &inventory_v1.Value_Int64Value{
					Int64Value: v,
				},
			}
		case float64:
			metadata[key] = &inventory_v1.Value{
				One: &inventory_v1.Value_DoubleValue{
					DoubleValue: v,
				},
			}
		case bool:
			metadata[key] = &inventory_v1.Value{
				One: &inventory_v1.Value_BoolValue{
					BoolValue: v,
				},
			}
		default:
			metadata[key] = nil
		}
	}

	return &inventory_v1.PartInfo{
		Uuid:        part.Uuid,
		Name:        part.Name,
		Description: part.Description,
		Price:       part.Price,
		Category:    StringToCategory(part.Category),
		Dimensions: &inventory_v1.Dimensions{
			Length: part.Dimensions.Length,
			Width:  part.Dimensions.Width,
			Height: part.Dimensions.Height,
			Weight: part.Dimensions.Weight,
		},
		Manufacturer: &inventory_v1.Manufacturer{
			Name:    part.Manufacturer.Name,
			Country: part.Manufacturer.Country,
			Website: part.Manufacturer.Website,
		},
		Tags:      part.Tags,
		Metadata:  metadata,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

func InventoryPartFilterToPartFilter(filter *inventory_v1.PartFilter) entity.Filter {
	var categories []string
	if filter == nil {
		return entity.Filter{}
	}
	for _, category := range filter.Categories {
		categories = append(categories, category.String())
	}
	return entity.Filter{
		Uuids:                 filter.Uuids,
		Names:                 filter.Names,
		Categories:            categories,
		ManufacturerCountries: filter.ManufacturerCountries,
		Tags:                  filter.Tags,
	}
}

func ArrayPartEntityToPartInfo(info []entity.Part) []*inventory_v1.PartInfo {
	convertedParts := make([]*inventory_v1.PartInfo, len(info))
	for i, _ := range info {
		convertedParts[i] = PartToInventoryPartInfo(info[i])
	}

	return convertedParts
}

func ArrayInventoryPartInfoToPartEntity(info []*inventory_v1.PartInfo) []entity.Part {
	convertedParts := make([]entity.Part, len(info))
	for i, _ := range info {
		convertedParts[i] = InventoryPartInfoToPart(info[i])
	}

	return convertedParts
}
