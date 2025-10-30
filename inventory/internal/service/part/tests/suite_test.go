package tests

import (
	"context"
	"github.com/Muvi7z/boilerplate/inventory/internal/repository/mock"
	"github.com/Muvi7z/boilerplate/inventory/internal/service"
	"github.com/Muvi7z/boilerplate/inventory/internal/service/part"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
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

	ctrl := gomock.NewController(s.T())

	s.partRepository = mock.NewMockPartRepository(ctrl)
	s.service = part.NewService(s.partRepository)
}

func (s *ServiceSuite) TearDownSuite() {}

func TestServiceIntegration(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}
