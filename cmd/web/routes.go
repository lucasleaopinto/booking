package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/lucasleaopinto/bookings/pkg/config"
	"github.com/lucasleaopinto/bookings/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {

	// mux := pat.New()
	// mux.Get("/",http.HandlerFunc(handlers.Repo.Home))
	// mux.Get("/about",http.HandlerFunc(handlers.Repo.About))

	mux := chi.NewRouter()
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/*", http.StripPrefix("/static",fileServer))

	return mux
}

//go get github.com/bmizerany/pat
