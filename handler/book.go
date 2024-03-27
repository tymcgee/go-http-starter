package handler

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/hlog"
	"github.com/tymcgee/go-starter/dao"
)

func (h *Handler) GetBooks(w http.ResponseWriter, r *http.Request) {
	l := hlog.FromRequest(r)
	q := dao.New(h.DB)

	books, err := q.GetBooks(r.Context())
	if err != nil {
		l.Error().Err(err).Msg("Failed to get books from database")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	out, err := json.Marshal(books)
	if err != nil {
		l.Error().Err(err).Any("books", books).Msg("Failed to marshal books into json")
	}
	l.Info().Msgf("Found %v books", len(books))
	w.Write(out)
}
