package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog/hlog"
	"github.com/rs/zerolog/log"
	"github.com/tymcgee/go-starter/config"
	"github.com/tymcgee/go-starter/handler"
	"github.com/tymcgee/go-starter/middleware"
)

func setupRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(hlog.NewHandler(log.Logger))
	r.Use(hlog.AccessHandler(func(r *http.Request, status, size int, duration time.Duration) {
		hlog.FromRequest(r).Info().
			Int("status", status).
			Int("size", size).
			Dur("duration", duration).
			Send()
	}))
	r.Use(hlog.RequestIDHandler("execution-id", "execution-id"))
	r.Use(hlog.RequestHandler("request"))
	// chi's recoverer has a nicer log output but in production we would want the structured log
	if config.Config.Environment == config.Local {
		r.Use(chiMiddleware.Recoverer)
	} else {
		r.Use(middleware.Recoverer)
	}

	h := Handler{}
	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Get("/health", h.Health)
		})
	})
	return r
}
