package api

import (
	"encoding/json"
	"net/http"

	"github.com/benspotatoes/extrawatch/models"

	"goji.io/pat"
)

func (rtr *Router) updatePlayer(w http.ResponseWriter, r *http.Request) {
	playerID := pat.Param(r, "player_id")

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

	err = rtr.Backend.UpdatePlayer(r.Context(), playerID, params)
	if err != nil {
		rtr.handleErrorResponse(w, r, http.StatusInternalServerError, err)
		return
	}

	player := &models.Player{ID: playerID}
	blob, err := json.Marshal(player)
	if err != nil {
		rtr.handleErrorResponse(w, r, http.StatusInternalServerError, err)
		return
	}

	rtr.handleSuccessResponse(w, r, http.StatusOK, blob)
}
