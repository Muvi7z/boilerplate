package app

import (
	"context"
	"errors"
	"fmt"
	"net"

	v1 "github.com/Muvi7z/boilerplate/iam/internal/api/iam/v1"
	"github.com/Muvi7z/boilerplate/iam/internal/config"
	"github.com/Muvi7z/boilerplate/iam/internal/migrator"
	"github.com/Muvi7z/boilerplate/iam/internal/repository/session"
	user2 "github.com/Muvi7z/boilerplate/iam/internal/repository/user"
	"github.com/Muvi7z/boilerplate/iam/internal/service/iam"
	"github.com/Muvi7z/boilerplate/iam/internal/service/user"
	"github.com/Muvi7z/boilerplate/platform/cache"
	redis "github.com/Muvi7z/boilerplate/platform/cache/redis"
	"github.com/Muvi7z/boilerplate/platform/closer"
	"github.com/Muvi7z/boilerplate/platform/grpc/health"
	"github.com/Muvi7z/boilerplate/platform/logger"
	iam_v1 "github.com/Muvi7z/boilerplate/shared/pkg/proto/iam/v1"
	redis2 "github.com/gomodule/redigo/redis"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" //
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type App struct {
	grpcServer *grpc.Server
	listener   net.Listener

	userService       *user.Service
	userRepository    *user2.Repository
	iamService        *iam.Service
	sessionRepository *session.Repository

	redisPool   *redis2.Pool
	redisClient cache.RedisClient

	db       *sqlx.DB
	migrator *migrator.Migrator
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

func (a *App) runGRPCServer(ctx context.Context) error {
	logger.Info(ctx, fmt.Sprintf("gRPC server listening on %s", config.AppConfig().IAMGRPC.Address()))

	err := a.grpcServer.Serve(a.listener)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initLogger,
		a.initListener,
		a.initCloser,
		a.InitIamGRPCServer,
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

func (a *App) initListener(_ context.Context) error {
	listener, err := net.Listen("tcp", config.AppConfig().IAMGRPC.Address())
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

func (a *App) InitIamGRPCServer(ctx context.Context) error {
	a.grpcServer = grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	closer.AddNamed("gRPC server", func(ctx context.Context) error {
		a.grpcServer.GracefulStop()
		return nil
	})

	reflection.Register(a.grpcServer)
	health.RegisterService(a.grpcServer)

	service, err := a.GetIamService()
	if err != nil {
		return err
	}

	apiIam := v1.NewApi(service)

	iam_v1.RegisterIAMServiceServer(a.grpcServer, apiIam)

	return nil
}

func (a *App) GetIamService() (*iam.Service, error) {
	if a.iamService == nil {
		userService, err := a.GetUserService()
		if err != nil {
			return nil, err
		}

		a.iamService = iam.New(
			a.getSessionRepository(),
			userService,
		)
	}

	return a.iamService, nil
}

func (a *App) GetUserService() (*user.Service, error) {
	if a.userService == nil {
		userRepository, err := a.getUserRepository()
		if err != nil {
			return nil, err
		}

		a.userService = user.New(userRepository)
	}

	return a.userService, nil
}

func (a *App) GetMigrator(ctx context.Context) (*migrator.Migrator, error) {
	if a.migrator == nil {
		db, err := a.GetDatabase()
		if err != nil {
			return nil, err
		}

		a.migrator = migrator.NewMigrator(db.DB, config.AppConfig().AppServerConfig.MigrationsDir())
	}

	return a.migrator, nil
}

func (a *App) GetDatabase() (*sqlx.DB, error) {
	if a.db == nil {
		uri := config.AppConfig().Postgres.URI()
		db, err := sqlx.Connect("postgres", uri)
		if err != nil {
			return nil, err
		}

		closer.AddNamed("Postgresql client", func(ctx context.Context) error {
			return db.Close()
		})

		a.db = db
	}

	return a.db, nil
}

func (a *App) getUserRepository() (*user2.Repository, error) {
	if a.userRepository == nil {
		db, err := a.GetDatabase()
		if err != nil {
			return nil, err
		}

		a.userRepository = user2.NewRepository(db)
	}

	return a.userRepository, nil
}

func (a *App) getSessionRepository() *session.Repository {
	if a.sessionRepository == nil {
		a.sessionRepository = session.NewRepository(a.getRedisClient())
	}

	return a.sessionRepository
}

func (a *App) getRedisClient() cache.RedisClient {
	if a.redisClient == nil {
		a.redisClient = redis.NewClient(
			a.getRedisPool(),
			logger.Logger(),
			config.AppConfig().Redis.ConnectionTimeout(),
		)
	}

	return a.redisClient
}

func (a *App) getRedisPool() *redis2.Pool {
	if a.redisPool == nil {
		a.redisPool = &redis2.Pool{
			MaxIdle:     config.AppConfig().Redis.MaxIdle(),
			IdleTimeout: config.AppConfig().Redis.IdleTimeout(),
			DialContext: func(ctx context.Context) (redis2.Conn, error) {
				return redis2.DialContext(ctx, "tcp", config.AppConfig().Redis.Address())
			},
		}
	}

	return a.redisPool
}
