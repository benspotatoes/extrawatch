package api

import (
	"encoding/json"
	"net/http"

	"goji.io/pat"
)

func (rtr *Router) getMatch(w http.ResponseWriter, r *http.Request) {
	matchID := pat.Param(r, "match_id")

	res, err := rtr.Backend.GetMatch(r.Context(), matchID)
	if err != nil {
		w.WriteHeader(http.StatusNotImplemented)
		return
	}

	blob, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(blob)
}
