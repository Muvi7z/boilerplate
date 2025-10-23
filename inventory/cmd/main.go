package main

import (
	"context"
	"fmt"
	inventoryv1 "github.com/Muvi7z/boilerplate/shared/pkg/proto/inventory/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
)

const grpcPort = 50051

type inventoryService struct {
	inventoryv1.UnimplementedInventoryServiceServer
}

func (s *inventoryService) GetPart(ctx context.Context, req *inventoryv1.GetPartRequest) (*inventoryv1.GetPartResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	part, ok := s.parts[req.GetUuid()]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "part %s not found", req.Uuid)
	}

	return &inventoryv1.GetPartResponse{
		Part: part,
	}, nil
}

func (s *inventoryService) ListPart(ctx context.Context, req *inventoryv1.ListPartRequest) (*inventoryv1.ListPartResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	filteredParts := make([]*inventoryv1.PartInfo, len(s.parts), 0)

	for i, _ := range s.parts {
		accepted := false
		index := 0

		filter := req.GetFilter()

		for _, fUuid := range filter.Uuids {
			if strings.Contains(fUuid, s.parts[i].Uuid) {
				accepted = true
			}
		}

		for _, fName := range filter.Names {
			if strings.Contains(fName, s.parts[i].Name) {
				accepted = true
			}
		}

		for _, fCategory := range filter.Categories {
			if strings.Contains(fCategory.String(), s.parts[i].Category.String()) {
				accepted = true
			}
		}

		for _, fManufacturerCountry := range filter.ManufacturerCountries {
			if strings.Contains(fManufacturerCountry, s.parts[i].Manufacturer.Country) {
				accepted = true
			}
		}

		for _, fTag := range filter.Tags {
			for _, tag := range s.parts[i].Tags {
				if strings.Contains(fTag, tag) {
					accepted = true
				}
			}

		}

		if len(filter.Uuids) == 0 && len(filter.Categories) == 0 && len(filter.ManufacturerCountries) == 0 &&
			len(filter.Names) == 0 && len(filter.Tags) == 0 {
			accepted = true
		}

		if accepted {
			filteredParts[index] = s.parts[i]
		}
	}
	return &inventoryv1.ListPartResponse{
		Parts: filteredParts,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return
	}

	defer func() {
		if err := lis.Close(); err != nil {
			log.Fatalf("failed to close listener: %v", err)
		}
	}()

	serverGrpc := grpc.NewServer()

	service := inventoryService{
		mu:    sync.RWMutex{},
		parts: make(map[string]*inventoryv1.PartInfo),
	}

	inventoryv1.RegisterInventoryServiceServer(serverGrpc, &service)

	reflection.Register(serverGrpc)

	go func() {
		err := serverGrpc.Serve(lis)
		if err != nil {
			log.Fatalf("failed to serve: %v", err)
			return
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("Shutting down server...")

	serverGrpc.GracefulStop()
	log.Println("Server gracefully stopped")

}
