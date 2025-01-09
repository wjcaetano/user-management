package router

import (
	"log"
	"net/http"
	"user-management/app/user/entrypoint/rest/handler"
	"user-management/internal/routes"

	"github.com/go-chi/chi/v5"
)

func NewRouter(handler *handler.Handler) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middlewareLogger)
	router.Use(middlewareRecoverer)

	routes.RegisterRoutes(router, handler)

	return router
}

func middlewareLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func middlewareRecoverer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic: %v", err)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
