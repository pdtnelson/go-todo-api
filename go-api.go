package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/joho/godotenv"
	"github.com/pdtnelson/go-api/models"
)

func routes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.RedirectSlashes,
		middleware.Recoverer,
	)

	router.Route("/v1", func(r chi.Router) {
		r.Mount("/api/todo", models.Routes(models.GetDB()))
	})

	return router
}

func main() {
	env := godotenv.Load()
	if env == nil {
		log.Fatal("Failed to load environment")
	} else {
		log.Print(env)
	}
	router := routes()
	log.Fatal(http.ListenAndServe(":8081", router))
}
