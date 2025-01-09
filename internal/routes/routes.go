package routes

import (
	"log"
	"net/http"
	"user-management/app/user/entrypoint/rest/handler"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(router *chi.Mux, handler *handler.Handler) {
	router.Get("/", homeHandler)
	router.Get("health", healthHandle)

	//handler.RegisterUserRoutes(router)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello World"))
	if err != nil {
		log.Println(err)
	}
}

func healthHandle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
