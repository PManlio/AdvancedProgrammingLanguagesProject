package middlewares

import (
	"net/http"

	"github.com/gorilla/handlers"
)

func SetCORS() func(http.Handler) http.Handler {
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	CORS := handlers.CORS(headers, methods, origins)
	return CORS
}
