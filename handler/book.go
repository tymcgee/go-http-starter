package handler

import (
	"net/http"

	"github.com/rs/zerolog/hlog"
	"github.com/tymcgee/go-http-starter/dao"
)

func (h Handler) GetBooks(w http.ResponseWriter, r *http.Request) {
	log := hlog.FromRequest(r)
	q := dao.New(h.DB)

	books, err := q.GetBooks(r.Context())
	if err != nil {
		log.Error().Err(err).Msg("Failed to get books from database")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Info().Msgf("Found %v books", len(books))
	RespondWithJson(w, books, log)
}
