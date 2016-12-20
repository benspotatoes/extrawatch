package api

import (
	"encoding/json"
	"net/http"

	"github.com/benspotatoes/extrawatch/models"
)

func (rtr *Router) postPlayerRound(w http.ResponseWriter, r *http.Request) {
	params := &models.PlayerRound{}
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

	err = rtr.Backend.InsertPlayerRound(r.Context(), params)
	if err != nil {
		rtr.handleErrorResponse(w, r, http.StatusInternalServerError, err)
		return
	}

	playerRound := &models.PlayerRound{PlayerID: params.PlayerID, RoundID: params.RoundID}
	blob, err := json.Marshal(playerRound)
	if err != nil {
		rtr.handleErrorResponse(w, r, http.StatusInternalServerError, err)
		return
	}

	rtr.handleSuccessResponse(w, r, http.StatusCreated, blob)
}
