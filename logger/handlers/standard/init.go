package standard

import (
	"log/slog"
	"os"

	"github.com/dhanar-kusuma/go-spark/environment"
	"github.com/dhanar-kusuma/go-spark/logger/handlers"
)

func Init(appName string, env environment.Type, opts ...any) (slog.Handler, handlers.Flush, error) {
	var slogHandlerOpts *slog.HandlerOptions
	var handler slog.Handler

	if len(opts) > 0 {
		if opt, ok := opts[0].(*slog.HandlerOptions); ok {
			slogHandlerOpts = opt
		} else {
			return nil, nil, ErrInvalidSlogHandlerOption
		}
	}
	if env == environment.Production || env == environment.Staging {
		handler = slog.NewJSONHandler(os.Stdout, slogHandlerOpts)
	} else {
		handler = slog.NewTextHandler(os.Stdout, slogHandlerOpts)
	}

	return handler.WithAttrs([]slog.Attr{
		{Key: handlers.LogAppName, Value: slog.StringValue(appName)},
		{Key: handlers.LogEnv, Value: slog.StringValue(env.String())},
	}), handlers.Void, nil
}
