package tests

import (
	"context"
	"github.com/Muvi7z/boilerplate/inventory/internal/repository/mock"
	"github.com/Muvi7z/boilerplate/inventory/internal/service"
	"github.com/Muvi7z/boilerplate/inventory/internal/service/part"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ServiceSuite struct {
	suite.Suite

	ctx context.Context

	partRepository *mock.MockPartRepository

	service service.PartService
}

func (s *ServiceSuite) SetupSuite() {
	s.ctx = context.Background()
	s.partRepository = mock.NewMockPartRepository(s.T())

	s.service = part.NewService(s.partRepository)
}

func (s *ServiceSuite) TearDownSuite() {}

func TestServiceIntegration(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}
