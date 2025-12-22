package integration

import (
	"context"
	"github.com/Muvi7z/boilerplate/platform/logger"
	"github.com/Muvi7z/boilerplate/platform/testcontainers"
	"github.com/Muvi7z/boilerplate/platform/testcontainers/network"
	"github.com/Muvi7z/boilerplate/platform/testcontainers/path"
	"github.com/Muvi7z/boilerplate/platform/testcontainers/postgres"
	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.uber.org/zap"
	"os"
	"time"
)

const (
	startupTimeout = 3 * time.Minute

	loggerLevelValue = "debug"
	appPortKey       = "APP_PORT"
	orderDockerfile  = "deploy"
)

type TestEnvironment struct {
	Network  *network.Network
	Postgres *postgres.Container
}

func setupTestEnvironment(ctx context.Context) *TestEnvironment {
	logger.Info(ctx, "Test environment setup")

	generatedNetwork, err := network.NewNetwork(ctx, projectName)
	if err != nil {
		logger.Fatal(ctx, "не удалось создать общую сеть")
	}

	logger.Info(ctx, "Сеть успешно создана")

	postgresUsername := getEnvWithLogging(ctx, testcontainers.PostgresUserKey)
	postgresPassword := getEnvWithLogging(ctx, testcontainers.PostgresPasswordKey)
	//postgresHost := getEnvWithLogging(ctx, testcontainers.PostgresHostKey)
	//postgresPort := getEnvWithLogging(ctx, testcontainers.PostgresPortKey)
	postgresDB := getEnvWithLogging(ctx, testcontainers.PostgresDBKey)
	postgresImage := getEnvWithLogging(ctx, testcontainers.PostgresImageNameKey)
	postgresSslmode := getEnvWithLogging(ctx, testcontainers.PostgresImageNameKey)

	appPort := getEnvWithLogging(ctx, appPortKey)

	generatedPostgres, err := postgres.NewContainer(ctx,
		postgres.WithNetworkName(generatedNetwork.Name()),
		postgres.WithContainerName(testcontainers.PostgresContainerName),
		postgres.WithImageName(postgresImage),
		postgres.WithDatabase(postgresDB),
		postgres.WithAuth(postgresUsername, postgresPassword),
		postgres.WithLogger(logger.Logger()),
		postgres.WithSslMode(postgresSslmode),
	)

	if err != nil {
		cleanupTestEnvironment(ctx, &TestEnvironment{Network: generatedNetwork})
		logger.Fatal(ctx, "не удалось запустить контейнер Postgres", zap.Error(err))
	}

	logger.Info(ctx, "✅ Контейнер Postgres успешно запущен")

	projectRoot := path.GetProjectRoot()

	appEnv := map[string]string{
		testcontainers.PostgresHostKey: generatedPostgres.Config().ContainerName,
	}

	waitStrategy := wait.ForListeningPort(nat.Port(appPort + "/tcp")).
		WithStartupTimeout(startupTimeout)

	_ = projectRoot
	_ = waitStrategy
	_ = appEnv

	//appContainer, err := app.NewContainer(ctx,
	//	app.WithName(ufoAppName),
	//	app.WithPort(grpcPort),
	//	app.WithDockerfile(projectRoot, ufoDockerfile),
	//	app.WithNetwork(generatedNetwork.Name()),
	//	app.WithEnv(appEnv),
	//	app.WithLogOutput(os.Stdout),
	//	app.WithStartupWait(waitStrategy),
	//	app.WithLogger(logger.Logger()),
	//)
	//if err != nil {
	//	cleanupTestEnvironment(ctx, &TestEnvironment{Network: generatedNetwork, Postgres: generatedPostgres})
	//	logger.Fatal(ctx, "не удалось запустить контейнер приложения", zap.Error(err))
	//}

	logger.Info(ctx, "🎉 Тестовое окружение готово")
	return &TestEnvironment{
		Network:  generatedNetwork,
		Postgres: generatedPostgres,
	}
}

// getEnvWithLogging возвращает значение переменной окружения с логированием
func getEnvWithLogging(ctx context.Context, key string) string {
	value := os.Getenv(key)
	if value == "" {
		logger.Warn(ctx, "Переменная окружения не установлена", zap.String("key", key))
	}

	return value
}
