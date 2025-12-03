package main

import (
	"context"
	"fmt"
	grpcinventory "github.com/Muvi7z/boilerplate/order/internal/client/grpc/inventory/v1"
	grpcpayment "github.com/Muvi7z/boilerplate/order/internal/client/grpc/payment/v1"
	orderhandler "github.com/Muvi7z/boilerplate/order/internal/handler/order"
	"github.com/Muvi7z/boilerplate/order/internal/migrator"
	"github.com/Muvi7z/boilerplate/order/internal/repository"
	"github.com/Muvi7z/boilerplate/order/internal/server"
	"github.com/Muvi7z/boilerplate/order/internal/usecase/order"
	inventory_v1 "github.com/Muvi7z/boilerplate/shared/pkg/proto/inventory/v1"
	payment_v1 "github.com/Muvi7z/boilerplate/shared/pkg/proto/payment/v1"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"os/signal"
	"syscall"
)

const serverAddress = "localhost:50052"
const grpcAddress = "localhost:50051"

func main() {
	conn, err := grpc.NewClient(
		grpcAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
		return
	}

	defer func() {
		if cerr := conn.Close(); cerr != nil {
			log.Fatal("failed to close connection to server:", cerr)
		}
	}()

	err = godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
		return
	}

	dbURI := os.Getenv("DB_URI")

	ctx := context.Background()

	pool, err := pgxpool.New(ctx, dbURI)
	if err != nil {
		log.Printf("failed to connect to database: %v\n", err)
		return
	}
	defer pool.Close()

	migrationsDir := os.Getenv("MIGRATIONS_DIR")

	migratorRunner := migrator.NewMigrator(stdlib.OpenDB(*pool.Config().ConnConfig.Copy()), migrationsDir)

	err = migratorRunner.Up()
	if err != nil {
		log.Printf("failed to run migrator: %v", err)
		return
	}

	paymentClient := payment_v1.NewPaymentClient(conn)
	inventoryClient := inventory_v1.NewInventoryServiceClient(conn)

	orderRepository := repository.New(pool.)

	grpcPaymentClient := grpcpayment.New(paymentClient)
	grpcInventoryClient := grpcinventory.New(inventoryClient)

	orderUsecase := order.New(grpcPaymentClient, grpcInventoryClient, orderRepository)

	orderHandler := orderhandler.NewHandler(orderUsecase)

	serverOpenApi := server.NewServer(orderHandler, fmt.Sprintf("%s", serverAddress))

	go func() {
		serverOpenApi.Run()
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("🛑 Shutting down Order server...")
	err = serverOpenApi.Shutdown(ctx)
	if err != nil {
		return
	}
	log.Println("✅ Server stopped")

}
