package main

import (
	"context"
	"fmt"
	apiPartV1 "github.com/Muvi7z/boilerplate/inventory/internal/api/inventory/v1"
	repoPart "github.com/Muvi7z/boilerplate/inventory/internal/repository/part"
	"github.com/Muvi7z/boilerplate/inventory/internal/service/part"
	inventoryv1 "github.com/Muvi7z/boilerplate/shared/pkg/proto/inventory/v1"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

const grpcPort = 50051

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
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

	err = godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
		return
	}

	dbURI := os.Getenv("MONGO_URI")
	dbName := os.Getenv("MONGO_INITDB_DATABASE")

	ctx := context.Background()

	client, err := mongo.Connect(options.Client().ApplyURI(dbURI))
	if err != nil {
		log.Printf("failed to connect to mongodb: %v\n", err)
		return
	}

	defer func() {
		cerr := client.Disconnect(ctx)
		if cerr != nil {
			log.Printf("failed to disconnect from mongodb: %v\n", cerr)
		}
	}()

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Printf("failed to ping mongodb: %v\n", err)
		return
	}

	db := client.Database(dbName)

	repositoryPart := repoPart.NewRepository(db)
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
