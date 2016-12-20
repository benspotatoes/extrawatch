package api

import (
	"net/http"

	"goji.io/pat"
)

func (rtr *Router) deletePlayerRound(w http.ResponseWriter, r *http.Request) {
	playerID := pat.Param(r, "player_id")
	roundID := pat.Param(r, "round_id")

	err := rtr.Backend.DeletePlayerRound(r.Context(), playerID, roundID)
	if err != nil {
		rtr.handleErrorResponse(w, r, http.StatusInternalServerError, err)
		return
	}

	rtr.handleSuccessResponse(w, r, http.StatusNoContent, nil)
}
