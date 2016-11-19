package api

import (
	"encoding/json"
	"net/http"

	"github.com/benspotatoes/extrawatch/models"
)

func (rtr *Router) postRound(w http.ResponseWriter, r *http.Request) {
	params := &models.Round{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		rtr.handleErrorResponse(w, r, http.StatusNotImplemented, err)
		return
	}

	err = params.Validate()
	if err != nil {
		rtr.handleErrorResponse(w, r, http.StatusBadRequest, err)
		return
	}

	id, err := rtr.Backend.InsertRound(r.Context(), params)
	if err != nil {
		rtr.handleErrorResponse(w, r, http.StatusInternalServerError, err)
		return
	}

	round := &models.Round{ID: id}
	blob, err := json.Marshal(round)
	if err != nil {
		rtr.handleErrorResponse(w, r, http.StatusInternalServerError, err)
		return
	}

	rtr.handleSuccessResponse(w, r, http.StatusCreated, blob)
}
