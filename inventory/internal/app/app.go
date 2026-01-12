package app

import (
	"context"
	"errors"
	"fmt"
	v1 "github.com/Muvi7z/boilerplate/inventory/internal/api/inventory/v1"
	"github.com/Muvi7z/boilerplate/inventory/internal/config"
	"github.com/Muvi7z/boilerplate/inventory/internal/repository"
	part2 "github.com/Muvi7z/boilerplate/inventory/internal/repository/part"
	"github.com/Muvi7z/boilerplate/inventory/internal/service"
	"github.com/Muvi7z/boilerplate/inventory/internal/service/part"
	"github.com/Muvi7z/boilerplate/platform/closer"
	"github.com/Muvi7z/boilerplate/platform/grpc/health"
	"github.com/Muvi7z/boilerplate/platform/logger"
	inventoryv1 "github.com/Muvi7z/boilerplate/shared/pkg/proto/inventory/v1"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"net"
)

type App struct {
	grpcServer     *grpc.Server
	listener       net.Listener
	partService    service.PartService
	partRepository repository.PartRepository

	mongoDBClient *mongo.Client
	mongoDBHandle *mongo.Database
}

func New(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run(ctx context.Context) error {
	return a.runGRPCServer(ctx)
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initLogger,
		a.initListener,
		a.initCloser,
		a.initGrpcServer,
	}

	for _, init := range inits {
		if err := init(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initLogger(_ context.Context) error {
	return logger.Init(
		config.AppConfig().Logger.Level(),
		config.AppConfig().Logger.AsJson(),
	)
}

func (a *App) getMongoDBClient(ctx context.Context) (*mongo.Client, error) {
	if a.mongoDBClient == nil {
		client, err := mongo.Connect(options.Client().ApplyURI(config.AppConfig().Mongo.URI()))
		if err != nil {
			panic(fmt.Sprintf("failed to connect to mongoDB: %s \n", err))
		}

		err = client.Ping(ctx, readpref.Primary())
		if err != nil {
			panic(fmt.Sprintf("failed to ping mongoDB: %s \n", err))
		}

		closer.AddNamed("MongoDB client", func(ctx context.Context) error {
			return client.Disconnect(ctx)
		})

		a.mongoDBClient = client
	}

	return a.mongoDBClient, nil
}

func (a *App) getMongoDBHandle(ctx context.Context) (*mongo.Database, error) {
	if a.mongoDBHandle == nil {
		client, err := a.getMongoDBClient(ctx)
		if err != nil {
			return nil, err
		}

		a.mongoDBHandle = client.Database(config.AppConfig().Mongo.DatabaseName())
	}

	return a.mongoDBHandle, nil
}

func (a *App) getPartRepository(ctx context.Context) repository.PartRepository {
	if a.partRepository == nil {

		mongoDBHandle, err := a.getMongoDBHandle(ctx)
		if err != nil {
			panic(err)
		}
		a.partRepository = part2.NewRepository(mongoDBHandle)
	}

	return a.partRepository
}

func (a *App) getPartService(ctx context.Context) service.PartService {
	if a.partService == nil {
		a.partService = part.NewService(a.getPartRepository(ctx))
	}

	return a.partService
}

func (a *App) initGrpcServer(ctx context.Context) error {
	a.grpcServer = grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	closer.AddNamed("gRPC server", func(ctx context.Context) error {
		a.grpcServer.GracefulStop()
		return nil
	})

	reflection.Register(a.grpcServer)

	health.RegisterService(a.grpcServer)

	apiPart := v1.NewAPI(a.getPartService(ctx))

	inventoryv1.RegisterInventoryServiceServer(a.grpcServer, apiPart)

	return nil
}

func (a *App) initListener(_ context.Context) error {
	listener, err := net.Listen("tcp", config.AppConfig().InventoryGRPC.Address())
	if err != nil {
		return err
	}

	closer.AddNamed("TCP listener", func(ctx context.Context) error {
		lerr := listener.Close()
		if lerr != nil && !errors.Is(lerr, net.ErrClosed) {
			return lerr
		}

		return nil
	})

	a.listener = listener

	return nil
}

func (a *App) initCloser(_ context.Context) error {
	closer.SetLogger(logger.Logger())

	return nil
}

func (a *App) runGRPCServer(ctx context.Context) error {
	logger.Info(ctx, fmt.Sprintf("gRPC server listening on %s", config.AppConfig().InventoryGRPC.Address()))

	err := a.grpcServer.Serve(a.listener)
	if err != nil {
		return err
	}

	return nil
}
