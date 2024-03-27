package middleware

import (
	"net/http"
	"runtime/debug"

	"github.com/rs/zerolog/hlog"
)

// Rolling my own recoverer to log the stack trace
// using zerolog. Based mostly on Chi's recoverer middleware.
func Recoverer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rvr := recover(); rvr != nil {
				if rvr == http.ErrAbortHandler {
					// we don't recover http.ErrAbortHandler so the response
					// to the client is aborted, this should not be logged
					panic(rvr)
				}

				// log the error
				log := hlog.FromRequest(r)
				log.Error().Bytes("stack", debug.Stack()).Msg("Panic during execution")

				if r.Header.Get("Connection") != "Upgrade" {
					w.WriteHeader(http.StatusInternalServerError)
				}
			}
		}()

		next.ServeHTTP(w, r)
	})
}
