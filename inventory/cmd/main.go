package main

import (
	"fmt"
	apiPartV1 "github.com/Muvi7z/boilerplate/inventory/internal/api/inventory/v1"
	repoPart "github.com/Muvi7z/boilerplate/inventory/internal/repository/part"
	"github.com/Muvi7z/boilerplate/inventory/internal/service/part"
	inventoryv1 "github.com/Muvi7z/boilerplate/shared/pkg/proto/inventory/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

const grpcPort = 50051

type inventoryService struct {
	inventoryv1.UnimplementedInventoryServiceServer
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return
	}

	defer func() {
		if cerr := lis.Close(); cerr != nil {
			log.Fatalf("failed to close listener: %v", err)
		}
	}()

	serverGrpc := grpc.NewServer()

	repositoryPart := repoPart.NewRepository()
	service := part.NewService(repositoryPart)
	apiPart := apiPartV1.NewAPI(service)

	inventoryv1.RegisterInventoryServiceServer(serverGrpc, apiPart)

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
