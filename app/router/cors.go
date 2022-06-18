package router

import (
	"github.com/gorilla/handlers"
	"net/http"
)

func WithCORS(handler http.Handler) http.Handler {
	methods := handlers.AllowedMethods([]string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete})
	headers := handlers.AllowedHeaders([]string{"Content-Type", "Accept", "X-Requested-with"})
	origins := handlers.AllowedOrigins([]string{"*"})
	return handlers.CORS(methods, headers, origins)(handler)
}
