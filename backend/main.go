package main

import (
	"net/http"
	"time"

	"./handler"
	"./waiter"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Logger)

	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(render.SetContentType(render.ContentTypeJSON))

	var reservations handler.ListOfReservations
	var tables waiter.ListOfTables

	r.Route("/reservations", func(r chi.Router) {
		r.Post("/", reservations.NewReservation)
		r.Get("/", reservations.GetReservations)

		r.Route("/{id}", func(r chi.Router) {
			r.Put("/", reservations.CancelReservation)
			r.Delete("/", reservations.DeleteReservation)
		})
	})

	r.Get("/tables", tables.GetTables)

	http.ListenAndServe(":3000", r)
}
