package e2e

import (
	"context"
	"github.com/Muvi7z/boilerplate/platform/logger"
	"go.uber.org/zap"
)

func teardownTestEnvironment(ctx context.Context, env *TestEnvironment) {
	log := logger.Logger()
	log.Info(ctx, "🧹 Очистка тестового окружения...")

	cleanupTestEnvironment(ctx, env)

	log.Info(ctx, "✅ Тестовое окружение успешно очищено")
}

// cleanupTestEnvironment — вспомогательная функция для освобождения ресурсов
func cleanupTestEnvironment(ctx context.Context, env *TestEnvironment) {
	if env.Postgres != nil {
		if err := env.Postgres.Terminate(ctx); err != nil {
			logger.Error(ctx, "не удалось остановить контейнер Postgres", zap.Error(err))
		} else {
			logger.Info(ctx, "🛑 Контейнер Postgres остановлен")
		}
	}

	if env.Network != nil {
		if err := env.Network.Remove(ctx); err != nil {
			logger.Error(ctx, "не удалось удалить сеть", zap.Error(err))
		} else {
			logger.Info(ctx, "Сеть удалена")
		}
	}
}
