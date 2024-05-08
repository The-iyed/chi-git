package api

import (
    "net/http"
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
	"go-chi/database"
)

func Run() {
    r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
    r.Get("/health-check", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("health check!"))
    })
	database.Connect()
    http.ListenAndServe(":5003", r)
}
