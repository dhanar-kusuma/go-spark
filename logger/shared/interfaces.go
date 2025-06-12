package shared

import (
	"log/slog"

	"github.com/dhanar-kusuma/go-spark/environment"
)

type Flush func() error

var Void Flush = func() error { return nil }

type LoggerFactory func(appName string, env environment.Type, opts ...any) (slog.Handler, Flush, error)
