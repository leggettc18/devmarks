package log

import (
	"context"
	"net/http"

	"github.com/sirupsen/logrus"
)

type Logger interface {
	WithError(error) *logrus.Entry
}

type contextKey struct { key string }

var loggerKey = contextKey{"logger"}

func GetLogger(ctx context.Context) *logrus.Logger{
	logger, ok := ctx.Value(loggerKey).(*logrus.Logger)
	if !ok {
		return nil
	}
	return logger
}

func setLoggerInCtx(ctx *context.Context, logger *logrus.Logger) {
	*ctx = context.WithValue(*ctx, loggerKey, logger)
}

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		logger := logrus.New()
		setLoggerInCtx(&ctx, logger)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}