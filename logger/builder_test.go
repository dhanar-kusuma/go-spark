package logger_test

import (
	"context"
	"testing"

	"github.com/dhanar-kusuma/go-spark/environment"
	"github.com/dhanar-kusuma/go-spark/logger"
	"github.com/stretchr/testify/assert"
)

func TestBuilder(t *testing.T) {
	appName := "go-spark"

	t.Run("no error with zap log", func(t *testing.T) {
		slogger, flush, err := logger.NewBuilder().
			SetEnv(environment.Production).
			SetLoggerType(logger.ZAP).
			Build(appName)
		assert.NoError(t, err)
		defer flush()

		rqLogger := slogger.With("request_id", "random_uuid")
		assert.NotNil(t, rqLogger)

		slogger.Info("example log message", "attr", "attr_value")
		slogger.InfoContext(context.Background(), "example log with context", "attr", "attr_value")

		rqLogger.Info("example log with request_id", "additional_attr", "attr_value_new")
	})
}
