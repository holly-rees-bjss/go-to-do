package middleware

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
)

type IDkey int

const TraceIDkey IDkey = 1

func TraceIDMiddleware(logger *slog.Logger, next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		traceID := uuid.New().String()
		ctx := context.WithValue(r.Context(), TraceIDkey, traceID)
		loggerWithCtx := logger.With("TraceID", traceID)
		slog.SetDefault(loggerWithCtx)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
