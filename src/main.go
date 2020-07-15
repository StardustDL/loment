package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"loment/handlers"
)

func main() {
	port := 80

	{
		portEnv := os.Getenv("LOMENT_PORT")
		p, err := strconv.Atoi(portEnv)
		if err == nil {
			port = p
		}
	}

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/", func(r chi.Router) {
		r.Get("/", handlers.All)
		r.Post("/", handlers.Create)
		r.Post("/query", handlers.Query)
		r.Route("/{id}", func(r chi.Router) {
			r.Use(handlers.ParamID)
			r.Get("/", handlers.Get)
			r.Delete("/", handlers.Delete)
			r.Put("/", handlers.Update)
		})
	})

	http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}
