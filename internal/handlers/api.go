package handlers

import (
	"github.com/go-chi/chi/v5"
	chimiddle "github.com/go-chi/chi/v5/middleware"
)

func Handler(r *chi.Mux) {
	//global middleware
	r.Use(chimiddle.StripSlashes)
}
