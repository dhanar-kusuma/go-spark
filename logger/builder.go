package logger

import (
	"log/slog"

	"github.com/dhanar-kusuma/go-spark/environment"
	"github.com/dhanar-kusuma/go-spark/logger/shared"
	"github.com/dhanar-kusuma/go-spark/logger/standard"
	"github.com/dhanar-kusuma/go-spark/logger/zap"
)

var factories = map[Type]shared.LoggerFactory{
	Standard: standard.Init,
	ZAP:      zap.Init,
}

type Builder struct {
	env        environment.Type
	loggerType Type
	options    []any
}

func NewBuilder() *Builder {
	return &Builder{
		env:        environment.Development,
		loggerType: Standard,
		options:    []any{},
	}
}

func (b *Builder) SetEnv(env environment.Type) *Builder {
	b.env = env
	return b
}

func (b *Builder) SetLoggerType(loggerType Type) *Builder {
	b.loggerType = loggerType
	return b
}

func (b *Builder) SetOptions(opts ...any) *Builder {
	b.options = opts
	return b
}

func (b *Builder) Build(appName string) (*slog.Logger, shared.Flush, error) {
	factory, found := factories[b.loggerType]
	if !found {
		return nil, nil, ErrUnsupportedLoggerType
	}
	handler, flush, err := factory(appName, b.env, b.options)
	if err != nil {
		return nil, nil, err
	}

	logger := slog.New(handler)
	return logger, flush, nil
}
