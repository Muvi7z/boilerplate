package tests

import (
	"github.com/Muvi7z/boilerplate/inventory/internal/entity"
	"github.com/brianvoe/gofakeit/v7"
	"time"
)

func (s *ServiceSuite) TestGetPartSuccess() {
	var (
		createdAt   = time.Now()
		updatedAt   = time.Now()
		uuid        = gofakeit.UUID()
		name        = gofakeit.Name()
		description = gofakeit.ProductDescription()
		price       = gofakeit.Price(0, 10000)

		partInfo = entity.Part{
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
		}
	)

	s.partRepository.EXPECT().
		GetPart(s.ctx, uuid)

	part, err := s.partRepository.GetPart(s.ctx, uuid)
	s.NoError(err)
	s.Equal(&partInfo, part)
}
