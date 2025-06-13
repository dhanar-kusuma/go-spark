package zap

import (
	"log/slog"

	"github.com/dhanar-kusuma/go-spark/environment"
	"github.com/dhanar-kusuma/go-spark/logger/handlers"
	zapPkg "go.uber.org/zap"
	"go.uber.org/zap/exp/zapslog"
)

func Init(appName string, env environment.Type, opts ...any) (slog.Handler, handlers.Flush, error) {
	var zapLogger *zapPkg.Logger
	var err error
	var zapOpts []zapPkg.Option

	if len(opts) > 0 {
		for _, v := range opts {
			if opt, ok := v.(zapPkg.Option); ok {
				zapOpts = append(zapOpts, opt)
			} else {
				return nil, nil, ErrInvalidZapOption
			}
		}
	}

	if env == environment.Production || env == environment.Staging {
		zapLogger, err = zapPkg.NewProduction(zapOpts...)
	} else {
		zapLogger, err = zapPkg.NewDevelopment(zapOpts...)
	}
	if err != nil {
		return nil, nil, err
	}

	slogHandler := zapslog.NewHandler(zapLogger.Core().With([]zapPkg.Field{
		zapPkg.String(handlers.LogAppName, appName),
		zapPkg.String(handlers.LogEnv, env.String()),
	}))
	return slogHandler, zapLogger.Sync, nil
}
