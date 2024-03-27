package main

import (
	"fmt"
	"net/http"

	"github.com/rs/zerolog/hlog"
)

type Handler struct{}

func (h *Handler) health(w http.ResponseWriter, r *http.Request) {
	log := hlog.FromRequest(r)
	log.Debug().Msg("Health check requested")
	fmt.Fprint(w, "OK")
}
