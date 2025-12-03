package part

import (
	"context"
	"github.com/Muvi7z/boilerplate/inventory/internal/entity"
	"github.com/Muvi7z/boilerplate/inventory/internal/repository/converter"
	repo "github.com/Muvi7z/boilerplate/inventory/internal/repository/entity"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *repository) ListPart(ctx context.Context, filter entity.Filter) ([]entity.Part, error) {
	filterDb := bson.M{}

	if len(filter.Uuids) > 0 {
		filterDb["_id"] = bson.M{"$in": filter.Uuids}
	}
	if len(filter.Names) > 0 {
		filterDb["name"] = bson.M{"$in": filter.Names}
	}
	if len(filter.Categories) > 0 {
		filterDb["category"] = bson.M{"$in": filter.Categories}
	}
	if len(filter.Tags) > 0 {
		filterDb["tags"] = bson.M{"$in": filter.Tags}
	}
	if len(filter.ManufacturerCountries) > 0 {
		filterDb["manufacturer.country"] = bson.M{
			"$in": filter.ManufacturerCountries,
		}
	}
	cursor, err := r.collection.Find(ctx, filterDb)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var parts []repo.Part

	if err = cursor.All(ctx, &parts); err != nil {
		return nil, err
	}

	return converter.ArrayRepoEntityToPartInfo(parts), nil

}
