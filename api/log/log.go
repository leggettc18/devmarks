package log

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"runtime/debug"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"leggett.dev/devmarks/api/model"
)

type Logger interface {
	WithError(error) *logrus.Entry
	LoggerMiddleware(http.Handler) http.Handler
}

type logSvc struct {
	*logrus.Logger
	ProxyCount int
}

func NewLogger(proxyCount int) Logger {
	return &logSvc{Logger: logrus.New(), ProxyCount: proxyCount}
}

type contextKey struct{ key string }

var loggerKey = contextKey{"logger"}

func GetLogger(ctx context.Context) *logrus.Logger {
	logger, ok := ctx.Value(loggerKey).(*logrus.Logger)
	if !ok {
		return nil
	}
	return logger
}

func setLoggerInCtx(ctx *context.Context, logger *logrus.Logger) {
	*ctx = context.WithValue(*ctx, loggerKey, logger)
}

type loggingResponseWriter struct {
	http.ResponseWriter
	status      int
	wroteHeader bool
}

func newLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{ResponseWriter: w}
}

func (lrw *loggingResponseWriter) Status() int {
	return lrw.status
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	if lrw.wroteHeader {
		return
	}
	lrw.status = code
	lrw.ResponseWriter.WriteHeader(code)
	lrw.wroteHeader = true
}

func (l *logSvc) ipAddressForRequest(r *http.Request) string {
	addr := r.RemoteAddr
	if l.ProxyCount > 0 {
		h := r.Header.Get("X-Forwarded-For")
		if h != "" {
			clients := strings.Split(h, ",")
			if l.ProxyCount > len(clients) {
				addr = clients[0]
			} else {
				addr = clients[len(clients)-l.ProxyCount]
			}
		}
	}
	//TODO: consider refactoring to use regex instead.
	if strings.Contains(addr, "[") { //If addr is ipv6
		sep_strings := strings.Split(strings.TrimSpace(addr), ":") //split string at the colons
		sep_strings = sep_strings[:len(sep_strings)-1]             //remove the last string (the port number)
		return strings.Join(sep_strings, ":")                      //Join the remaining elements back together into one string with the colons in between.
	} //If addr is ipv4
	return strings.Split(strings.TrimSpace(addr), ":")[0]
}

func (l *logSvc) LoggerMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				l.Error(fmt.Errorf("%v: %s", r, debug.Stack()))
			}
		}()

		start := time.Now()
		lrw := newLoggingResponseWriter(w)
		next.ServeHTTP(lrw, r)
		duration := time.Since(start)
		address := l.ipAddressForRequest(r)
		logger := l.WithFields(logrus.Fields{
			"request_id": base64.RawURLEncoding.EncodeToString(model.NewID()),
			"duration":    duration,
			"status_code": lrw.status,
			"remote":      address,
		})
		logger.Info(r.Method + " " + r.URL.RequestURI())
	}
	return http.HandlerFunc(fn)
}
