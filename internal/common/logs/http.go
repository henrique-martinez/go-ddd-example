package logs

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

func NewStructuredLogger() func(next http.Handler) http.Handler {
	return middleware.RequestLogger(&StructuredLogger{})
}

// based on example from chi: https://github.com/go-chi/chi/blob/master/_examples/logging/main.go
type StructuredLogger struct {
}

func (l *StructuredLogger) NewLogEntry(r *http.Request) middleware.LogEntry {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	sugar := logger.Sugar()

	if reqID := middleware.GetReqID(r.Context()); reqID != "" {
		sugar = sugar.With("req_id", reqID)
	}
	sugar = sugar.With(
		"http_method", r.Method,
		"remote_addr", r.RemoteAddr,
		"uri", r.RequestURI,
	)

	entry := &StructuredLoggerEntry{Logger: sugar}
	entry.Logger.Info("Request started")

	return entry
}

type StructuredLoggerEntry struct {
	Logger *zap.SugaredLogger
}

func (l *StructuredLoggerEntry) Write(status, bytes int, header http.Header, elapsed time.Duration, extra interface{}) {
	l.Logger = l.Logger.With(
		"resp_status", status,
		"resp_bytes_length", bytes,
		"resp_elapsed", elapsed.Round(time.Millisecond/100).String(),
	)

	l.Logger.Info("Request completed	")
}

func (l *StructuredLoggerEntry) Panic(v interface{}, stack []byte) {
	l.Logger = l.Logger.With(
		"stack", string(stack),
		"panic", fmt.Sprintf("%+v", v),
	)
}

func GetLogEntry(r *http.Request) *zap.SugaredLogger {
	entry := middleware.GetLogEntry(r).(*StructuredLoggerEntry)
	return entry.Logger
}
