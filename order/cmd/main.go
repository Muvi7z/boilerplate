package main

import (
	"context"
	"fmt"
	app2 "github.com/Muvi7z/boilerplate/order/internal/app"
	"github.com/Muvi7z/boilerplate/order/internal/config"
	"github.com/Muvi7z/boilerplate/platform/closer"
	"github.com/Muvi7z/boilerplate/platform/logger"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"os/signal"
	"syscall"
)

const serverAddress = "localhost:50052"
const grpcAddress = "localhost:50051"
const grpcPaymentAddress = "localhost:50053"
const configPath = "../../deploy/compose/order/.env"

func main() {
	err := config.Load(configPath)
	if err != nil {
		panic(fmt.Errorf("error loading config: %w", err))
	}

	appCtx, appCancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	defer appCancel()

	closer.Configure(syscall.SIGINT, syscall.SIGTERM)

	app, err := app2.New(appCtx)
	if err != nil {
		logger.Error(appCtx, "❌ Не удалось создать приложение", zap.Error(err))
		return
	}

	err = app.Run(appCtx)
	if err != nil {
		logger.Error(appCtx, "❌ Ошибка при работе приложения", zap.Error(err))
		return
	}
	//
	//conn, err := grpc.NewClient(
	//	grpcAddress,
	//	grpc.WithTransportCredentials(insecure.NewCredentials()),
	//)
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}
	//
	//defer func() {
	//	if cerr := conn.Close(); cerr != nil {
	//		log.Fatal("failed to close connection to server:", cerr)
	//	}
	//}()
	//
	//connPayment, err := grpc.NewClient(
	//	grpcPaymentAddress,
	//	grpc.WithTransportCredentials(insecure.NewCredentials()),
	//)
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}
	//
	//defer func() {
	//	if cerr := connPayment.Close(); cerr != nil {
	//		log.Fatal("failed to close connection to server:", cerr)
	//	}
	//}()
	//
	//err = godotenv.Load(".env")
	//if err != nil {
	//	log.Println("Error loading .env file")
	//	return
	//}
	//
	//dbURI := os.Getenv("DB_URI")
	//
	//host := os.Getenv("POSTGRES_HOST")
	//port := os.Getenv("POSTGRES_PORT")
	//user := os.Getenv("POSTGRES_USER")
	//password := os.Getenv("POSTGRES_PASSWORD")
	//db := os.Getenv("POSTGRES_DB")
	//sslmode := os.Getenv("POSTGRES_SSL_MODE")
	//
	//ctx := context.Background()
	//
	//pool, err := pgxpool.New(ctx, dbURI)
	//if err != nil {
	//	log.Printf("failed to connect to database: %v\n", err)
	//	return
	//}
	//defer pool.Close()
	//
	//migrationsDir := os.Getenv("MIGRATIONS_DIR")
	//
	//migratorRunner := migrator.NewMigrator(stdlib.OpenDB(*pool.Config().ConnConfig.Copy()), migrationsDir)
	//
	//err = migratorRunner.Up()
	//if err != nil {
	//	log.Printf("failed to run migrator: %v", err)
	//	return
	//}
	//
	//paymentClient := payment_v1.NewPaymentClient(connPayment)
	//inventoryClient := inventory_v1.NewInventoryServiceClient(conn)
	//
	//connectStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
	//	host,
	//	port,
	//	user,
	//	db,
	//	password,
	//	sslmode)
	//
	//connDb, err := sqlx.Connect("postgres", connectStr)
	//if err != nil {
	//	log.Printf("failed to connect to postgres: %v\n", err)
	//	return
	//}
	//
	//if err = connDb.Ping(); err != nil {
	//	log.Printf("cant ping db %v", err)
	//	return
	//}
	//
	//orderRepository := repository.New(connDb)
	//
	//grpcPaymentClient := grpcpayment.New(paymentClient)
	//grpcInventoryClient := grpcinventory.New(inventoryClient)
	//
	//orderUsecase := order.New(grpcPaymentClient, grpcInventoryClient, orderRepository)
	//
	//orderHandler := orderhandler.NewHandler(orderUsecase)
	//
	//serverOpenApi := server.NewServer(orderHandler, fmt.Sprintf("%s", serverAddress))
	//
	//go func() {
	//	serverOpenApi.Run()
	//}()
	//
	//// Graceful shutdown
	//quit := make(chan os.Signal, 1)
	//signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	//<-quit
	//log.Println("🛑 Shutting down Order server...")
	//err = serverOpenApi.Shutdown(ctx)
	//if err != nil {
	//	return
	//}
	//log.Println("✅ Server stopped")

}

func gracefulShutdown() {

}
