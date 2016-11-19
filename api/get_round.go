package api

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"goji.io/pat"
)

func (rtr *Router) getRound(w http.ResponseWriter, r *http.Request) {
	roundID := pat.Param(r, "round_id")

	round, err := rtr.Backend.SelectRound(r.Context(), roundID)
	if err == sql.ErrNoRows {
		rtr.handleErrorResponse(w, r, http.StatusNotFound, nil)
		return
	} else if err != nil {
		rtr.handleErrorResponse(w, r, http.StatusInternalServerError, err)
		return
	}

	blob, err := json.Marshal(round)
	if err != nil {
		rtr.handleErrorResponse(w, r, http.StatusInternalServerError, err)
		return
	}

	rtr.handleSuccessResponse(w, r, http.StatusOK, blob)
}
