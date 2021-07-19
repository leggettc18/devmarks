package log

import (
	"context"
	"encoding/base64"
	"net/http"

	"github.com/sirupsen/logrus"
	"leggett.dev/devmarks/api/model"
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

func setLoggerInCtx(ctx *context.Context, logger logrus.FieldLogger) {
	*ctx = context.WithValue(*ctx, loggerKey, logger)
}

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		logger := logrus.WithField("request_id", base64.RawURLEncoding.EncodeToString(model.NewID()))
		setLoggerInCtx(&ctx, logger)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}