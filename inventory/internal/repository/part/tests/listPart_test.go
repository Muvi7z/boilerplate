package tests

import (
	"github.com/Muvi7z/boilerplate/inventory/internal/entity"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/labstack/gommon/log"
	"time"
)

func (s *RepositorySuite) TestListPartSuccessFilterUUID() {
	var (
		createdAt   = time.Now()
		updatedAt   = time.Now()
		uuid        = gofakeit.UUID()
		uuid2       = gofakeit.UUID()
		name        = gofakeit.Name()
		description = gofakeit.ProductDescription()
		price       = gofakeit.Price(0, 10000)

		filter = entity.Filter{
			Uuids: []string{uuid, uuid2},
		}

		listPart = []entity.Part{
			{
				Uuid:        uuid,
				Name:        name,
				Description: description,
				Price:       price,
				Category:    entity.ENGINE,
				Dimensions: entity.Dimensions{
					Length: gofakeit.Float64(),
					Width:  gofakeit.Float64(),
					Height: gofakeit.Float64(),
					Weight: gofakeit.Float64(),
				},
				Manufacturer: entity.Manufacturer{
					Name:    gofakeit.Company(),
					Country: gofakeit.Country(),
					Website: gofakeit.Email(),
				},
				Tags:      nil,
				Metadata:  nil,
				CreatedAt: &createdAt,
				UpdatedAt: &updatedAt,
			},
			{
				Uuid:        gofakeit.UUID(),
				Name:        name,
				Description: description,
				Price:       price,
				Category:    entity.ENGINE,
				Dimensions: entity.Dimensions{
					Length: gofakeit.Float64(),
					Width:  gofakeit.Float64(),
					Height: gofakeit.Float64(),
					Weight: gofakeit.Float64(),
				},
				Manufacturer: entity.Manufacturer{
					Name:    gofakeit.Company(),
					Country: gofakeit.Country(),
					Website: gofakeit.Email(),
				},
				Tags:      nil,
				Metadata:  nil,
				CreatedAt: &createdAt,
				UpdatedAt: &updatedAt,
			},
			{
				Uuid:        uuid2,
				Name:        name,
				Description: description,
				Price:       price,
				Category:    entity.ENGINE,
				Dimensions: entity.Dimensions{
					Length: gofakeit.Float64(),
					Width:  gofakeit.Float64(),
					Height: gofakeit.Float64(),
					Weight: gofakeit.Float64(),
				},
				Manufacturer: entity.Manufacturer{
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
	)

	list, err := s.partRepository.ListPart(s.ctx, filter)
	log.Info(list)

	s.NoError(err)
	s.Equal(listPart, list)
}
