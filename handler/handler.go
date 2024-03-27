package handler

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/hlog"
)

type Handler struct {
	DB *sql.DB
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	log := hlog.FromRequest(r)
	log.Debug().Msg("Health check requested")
	fmt.Fprint(w, "OK")
}
