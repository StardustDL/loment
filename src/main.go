package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"

	"loment/handlers"
	"loment/repositories"
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

	dbSource := os.Getenv("LOMENT_DBORIGIN")
	dbName := os.Getenv("LOMENT_DBNAME")
	debug := os.Getenv("LOMENT_DEBUG")
	repo := repositories.Create(dbSource, dbName)

	err := repo.EnsureExisits()
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	err = repo.Start(debug != "")
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	defer repo.Stop()

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(cors.Handler(cors.Options{
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Use(middleware.Timeout(60 * time.Second))

	handler := new(handlers.CommentHandler)
	handler.Repo = repo

	r.Route("/", func(r chi.Router) {
		r.Post("/", handler.Create)     // comment -> id
		r.Post("/query", handler.Query) // commentQuery -> list[comment] or null
		r.Post("/count", handler.Count) // commentQuery -> number
		r.Route("/{id}", func(r chi.Router) {
			r.Use(handlers.ParamID)
			r.Get("/", handler.Get)       // id -> comment
			r.Delete("/", handler.Delete) // id -> bool
			r.Put("/", handler.Update)    // id, comment -> bool
		})
	})

	http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}
