package main

import (
	"context"
	"fmt"
	"github.com/Muvi7z/boilerplate/assembly/internal/app"
	"github.com/Muvi7z/boilerplate/assembly/internal/config"
	"github.com/Muvi7z/boilerplate/platform/closer"
	"github.com/Muvi7z/boilerplate/platform/logger"
	"go.uber.org/zap"
	"os/signal"
	"syscall"
)

const envPath = "./deploy/compose/assembly/.env"

func main() {
	err := config.Load(envPath)
	if err != nil {
		panic(fmt.Errorf("failed to load configuration: %w", err))
	}

	appCtx, appCancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer appCancel()

	closer.Configure(syscall.SIGINT, syscall.SIGTERM)

	a, err := app.New(appCtx)
	if err != nil {
		logger.Error(appCtx, "❌ Не удалось создать приложение", zap.Error(err))
		panic(err)
	}

	err = a.Run(appCtx)
	if err != nil {
		logger.Error(appCtx, "❌ Ошибка при работе приложения", zap.Error(err))
		panic(err)
	}
}
