package api

import (
	"net/http"

	"goji.io/pat"
)

func (rtr *Router) deleteMatch(w http.ResponseWriter, r *http.Request) {
	matchID := pat.Param(r, "match_id")

	err := rtr.Backend.DeleteMatch(r.Context(), matchID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
