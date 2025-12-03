package part

import (
	"context"
	"github.com/Muvi7z/boilerplate/inventory/internal/entity"
	repoEntity "github.com/Muvi7z/boilerplate/inventory/internal/repository/entity"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"time"
)

func Init(collection *mongo.Collection) {
	fUuid := uuid.New().String()

	createdAt := time.Now()
	updatedAt := time.Now()

	parts := []repoEntity.Part{
		{
			Uuid:        fUuid,
			Name:        gofakeit.Name(),
			Description: gofakeit.ProductDescription(),
			Price:       gofakeit.Price(100, 10000),
			Category:    entity.ENGINE,
			Dimensions: repoEntity.DimensionsInfo{
				Length: gofakeit.Float64(),
				Width:  gofakeit.Float64(),
				Height: gofakeit.Float64(),
				Weight: gofakeit.Float64(),
			},
			Manufacturer: repoEntity.ManufacturerInfo{
				Name:    gofakeit.Company(),
				Country: gofakeit.Country(),
				Website: gofakeit.Email(),
			},
			Tags:      nil,
			Metadata:  nil,
			CreatedAt: &createdAt,
			UpdatedAt: &updatedAt,
		},
	}

	_, err := collection.InsertMany(context.Background(), parts)
	if err != nil {
		return
	}

}
