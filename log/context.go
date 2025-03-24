package log

import (
	"context"

	log "github.com/sirupsen/logrus"
)

var LoggerCtxKey = "logger"

func GetLoggerFromContext(ctx context.Context) *log.Entry {
	logger := ctx.Value(LoggerCtxKey) //nolint:staticcheck
	if logger == nil {
		return log.NewEntry(log.StandardLogger())
	}

	entry, ok := logger.(*log.Entry)
	if !ok {
		return log.NewEntry(log.StandardLogger())
	}

	return entry
}

func WithLogger(ctx context.Context, logger *log.Entry) context.Context {
	return context.WithValue(ctx, LoggerCtxKey, logger)
}
