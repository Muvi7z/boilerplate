package part

import (
	"context"
	"errors"
	"github.com/Muvi7z/boilerplate/inventory/internal/entity"
	"github.com/Muvi7z/boilerplate/inventory/internal/repository/converter"
	repo "github.com/Muvi7z/boilerplate/inventory/internal/repository/entity"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (r *repository) GetPart(ctx context.Context, uuid string) (entity.Part, error) {
	var part repo.Part

	err := r.collection.FindOne(ctx, bson.M{"_id": uuid}).Decode(&part)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return entity.Part{}, entity.ErrPartInfoNotFound
		}

		return entity.Part{}, err
	}

	return converter.RepoEntityToPartInfo(part), nil
}
