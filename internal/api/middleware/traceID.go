package middleware

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
)

func TraceIDMiddleware(logger *slog.Logger, next http.Handler) http.Handler {
	type IDkey int
	const traceIDkey IDkey = 1

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		traceID := uuid.New().String()
		ctx := context.WithValue(r.Context(), traceIDkey, traceID)
		loggerWithCtx := logger.With("TraceID", traceID)
		ctx = context.WithValue(ctx, "logger", loggerWithCtx)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
