package tests

import (
	"context"
	"github.com/Muvi7z/boilerplate/inventory/internal/repository"
	"github.com/Muvi7z/boilerplate/inventory/internal/repository/part"
	"github.com/stretchr/testify/suite"
	"testing"
)

type RepositorySuite struct {
	suite.Suite

	ctx context.Context

	partRepository repository.PartRepository
}

func (s *RepositorySuite) SetupSuite() {
	s.ctx = context.Background()
	s.partRepository = part.NewRepository()

}

func (s *RepositorySuite) TearDownSuite() {
}

func TestResponseIntegration(t *testing.T) {
	suite.Run(t, new(RepositorySuite))
}
