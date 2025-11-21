package cmd

import (
	"context"
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
	"log"
	"os"
)

const serverAddress = "localhost:50051"

func main() {
	conn, err := grpc.NewClient(serverAddress)
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

	orderRepository := repository.New()

	grpcPaymentClient := grpcpayment.New(paymentClient)
	grpcInventoryClient := grpcinventory.New(inventoryClient)

	orderUsecase := order.New(grpcPaymentClient, grpcInventoryClient, orderRepository)

	orderHandler := orderhandler.NewHandler(orderUsecase)

	serverOpenApi := server.NewServer(orderHandler)

	go func() {
		serverOpenApi.Run()
	}()

}
