package e2e

import (
	"context"
	"github.com/Muvi7z/boilerplate/platform/logger"
	"github.com/Muvi7z/boilerplate/platform/testcontainers/network"
	"github.com/Muvi7z/boilerplate/platform/testcontainers/postgres"
	"go.uber.org/zap"
	"os"
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

	postgresUsername := getEnvWithLogging(ctx)

}

// getEnvWithLogging возвращает значение переменной окружения с логированием
func getEnvWithLogging(ctx context.Context, key string) string {
	value := os.Getenv(key)
	if value == "" {
		logger.Warn(ctx, "Переменная окружения не установлена", zap.String("key", key))
	}

	return value
}
