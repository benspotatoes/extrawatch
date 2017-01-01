package api

import (
	"encoding/json"
	"net/http"

	"github.com/benspotatoes/extrawatch/models"
)

func (rtr *Router) postPlayer(w http.ResponseWriter, r *http.Request) {
	params := &models.Player{}
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

	id, err := rtr.Backend.InsertPlayer(r.Context(), params)
	if err != nil {
		rtr.handleErrorResponse(w, r, http.StatusInternalServerError, err)
		return
	}

	player := &models.Player{ID: id}
	blob, err := json.Marshal(player)
	if err != nil {
		rtr.handleErrorResponse(w, r, http.StatusInternalServerError, err)
		return
	}

	rtr.handleSuccessResponse(w, r, http.StatusCreated, blob)
}
