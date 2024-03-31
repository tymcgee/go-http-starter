package handler

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog"
)

// Marshals data into json and writes it to the ResponseWriter.
// Logs and responds with an error if marshaling to json fails.
func RespondWithJson(w http.ResponseWriter, data any, log *zerolog.Logger) {
	response, err := json.Marshal(data)
	if err != nil {
		if log == nil {
			log = zerolog.DefaultContextLogger
		}
		log.Error().Err(err).Msg("Failed to marshal response body into json")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(response)
}
