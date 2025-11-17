package cmd

import (
	grpcinventory "github.com/Muvi7z/boilerplate/order/internal/client/grpc/inventory/v1"
	grpcpayment "github.com/Muvi7z/boilerplate/order/internal/client/grpc/payment/v1"
	orderhandler "github.com/Muvi7z/boilerplate/order/internal/handler/order"
	"github.com/Muvi7z/boilerplate/order/internal/repository"
	"github.com/Muvi7z/boilerplate/order/internal/server"
	"github.com/Muvi7z/boilerplate/order/internal/usecase/order"
	inventory_v1 "github.com/Muvi7z/boilerplate/shared/pkg/proto/inventory/v1"
	payment_v1 "github.com/Muvi7z/boilerplate/shared/pkg/proto/payment/v1"
	"google.golang.org/grpc"
	"log"
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
