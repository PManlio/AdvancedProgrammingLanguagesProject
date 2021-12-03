package middlewares

import (
	"net/http"
)

func GlobalHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json") // "application/json"

		next.ServeHTTP(w, r)
	})
}
