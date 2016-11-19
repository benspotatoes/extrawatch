package api

import (
	"encoding/json"
	"net/http"

	"goji.io/pat"

	"github.com/benspotatoes/extrawatch/models"
)

func (rtr *Router) updateRound(w http.ResponseWriter, r *http.Request) {
	roundID := pat.Param(r, "round_id")

	params := &models.Round{}
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

	err = rtr.Backend.UpdateRound(r.Context(), roundID, params)
	if err != nil {
		rtr.handleErrorResponse(w, r, http.StatusInternalServerError, err)
		return
	}

	round := &models.Round{ID: roundID}
	blob, err := json.Marshal(round)
	if err != nil {
		rtr.handleErrorResponse(w, r, http.StatusInternalServerError, err)
		return
	}

	rtr.handleSuccessResponse(w, r, http.StatusOK, blob)
}
