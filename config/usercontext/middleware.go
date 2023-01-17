package usercontext

import (
	"context"
	"net/http"
)

func CorrelationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.Header.Get("X-Correlation-Id")
		ctx := context.WithValue(r.Context(), "X-Correlation-Id", id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
