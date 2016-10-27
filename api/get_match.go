package api

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"goji.io/pat"
)

func (rtr *Router) getMatch(w http.ResponseWriter, r *http.Request) {
	matchID := pat.Param(r, "match_id")

	match, err := rtr.Backend.SelectMatch(r.Context(), matchID)
	if err == sql.ErrNoRows {
		rtr.handleErrorResponse(w, r, http.StatusNotFound, nil)
		return
	} else if err != nil {
		rtr.handleErrorResponse(w, r, http.StatusInternalServerError, err)
		return
	}

	blob, err := json.Marshal(match)
	if err != nil {
		rtr.handleErrorResponse(w, r, http.StatusInternalServerError, err)
		return
	}

	rtr.handleSuccessResponse(w, r, http.StatusOK, blob)
}
