package zap

import (
	"context"
	"log/slog"

	"go.uber.org/zap"
)

type ZapHandler struct {
	logger *zap.Logger
}

func NewZapHandler(zapLogger *zap.Logger) *ZapHandler {
	return &ZapHandler{
		logger: zapLogger,
	}
}

func (z *ZapHandler) Enabled(_ context.Context, level slog.Level) bool {
	return true
}

func (z *ZapHandler) Handle(ctx context.Context, record slog.Record) error {
	zapFields := []zap.Field{}
	record.Attrs(func(a slog.Attr) bool {
		zapFields = append(zapFields, zap.Any(a.Key, a.Value))
		return true
	})

	msg := record.Message

	switch record.Level {
	case slog.LevelDebug:
		z.logger.Debug(msg, zapFields...)
	case slog.LevelInfo:
		z.logger.Info(msg, zapFields...)
	case slog.LevelWarn:
		z.logger.Warn(msg, zapFields...)
	case slog.LevelError:
		z.logger.Error(msg, zapFields...)
	default:
		z.logger.Info(msg, zapFields...)
	}
	return nil
}

func (z *ZapHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return z
}

func (z *ZapHandler) WithGroup(name string) slog.Handler {
	return z
}
