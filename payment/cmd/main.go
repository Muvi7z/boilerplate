package main

import (
	"context"
	"fmt"
	v1 "github.com/Muvi7z/boilerplate/payment/internal/api/payment/v1"
	"github.com/Muvi7z/boilerplate/payment/internal/service/payment"
	paymentv1 "github.com/Muvi7z/boilerplate/shared/pkg/proto/payment/v1"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

const grpcPort = 50053

type PaymentService struct {
	paymentv1.UnimplementedPaymentServer
}

func (s *PaymentService) PayOrder(ctx context.Context, req *paymentv1.PayOrderRequest) (*paymentv1.PayOrderResponse, error) {

	transactionUuid := uuid.New()

	log.Printf("Оплата прошла успешно, transaction_uuid: %s", transactionUuid.String())

	return &paymentv1.PayOrderResponse{
		TransactionUuid: transactionUuid.String(),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
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

	service := payment.NewService()

	api := v1.NewAPI(service)

	paymentv1.RegisterPaymentServer(serverGrpc, api)

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
