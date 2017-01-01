package api

import (
	"encoding/json"
	"net/http"

	"github.com/benspotatoes/extrawatch/models"

	"goji.io/pat"
)

func (rtr *Router) updateMatch(w http.ResponseWriter, r *http.Request) {
	matchID := pat.Param(r, "match_id")

	params := &models.Match{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		rtr.handleErrorResponse(w, r, http.StatusInternalServerError, err)
		return
	}

	err = params.Validate()
	if err != nil {
		rtr.handleErrorResponse(w, r, http.StatusBadRequest, err)
		return
	}

	err = rtr.Backend.UpdateMatch(r.Context(), matchID, params)
	if err != nil {
		rtr.handleErrorResponse(w, r, http.StatusInternalServerError, err)
		return
	}

	match := &models.Match{ID: matchID}
	blob, err := json.Marshal(match)
	if err != nil {
		rtr.handleErrorResponse(w, r, http.StatusInternalServerError, err)
		return
	}

	rtr.handleSuccessResponse(w, r, http.StatusOK, blob)
}
