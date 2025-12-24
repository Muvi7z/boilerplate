//go:build integration

package integration

import (
	"context"
	"fmt"
	"github.com/Muvi7z/boilerplate/platform/logger"
	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"
	"os"
	"path/filepath"
	"testing"
	"time"
)

const testsTimeout = 5 * time.Minute

var (
	env         *TestEnvironment
	suiteCtx    context.Context
	suiteCancel context.CancelFunc
)

func TestIntegration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Order Service Integration Test Suite")

}

var _ = BeforeSuite(func() {
	err := logger.Init(loggerLevelValue, true)
	if err != nil {
		panic(fmt.Sprintf("Failed to initialize logger: %v", err))
	}

	suiteCtx, suiteCancel = context.WithTimeout(context.Background(), testsTimeout)

	envVars, err := godotenv.Read(filepath.Join("..", "..", "..", "deploy", "compose", "order", ".env"))

	if err != nil {
		logger.Fatal(suiteCtx, "Failed to read env file", zap.Error(err))
	}

	for key, value := range envVars {
		_ = os.Setenv(key, value)
	}

	logger.Info(suiteCtx, "Запуск тестового окружения...")
	env = setupTestEnvironment(suiteCtx)
})
