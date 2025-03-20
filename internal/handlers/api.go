package handlers

import (
	"github.com/TheAmgadX/APIs-With-GoLang/internal/middleware"
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux) {
	/*
		Global Middleware -> will be used on all endpoints
		we make (applied all the time)

		Middleware is a function that runs before the actual route handler
		Usually used for security, logging, modifying requests, etc.
	*/
	r.Use(chimiddle.StripSlashes) /* to ignore strip slashes like that :
	http://localhost:8080/user/ --> without StripSlashes it will cause: 404 Not Found

	router.Use() used to apply middleware functions to routes
	*/

	r.Route("/account", func(router chi.Router) {
		// middleware for /account route not global.
		router.Use(middleware.Authorization)

		router.Get("/coins", GetCoinBalance)
	})

}
