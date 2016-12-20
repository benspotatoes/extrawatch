package api

import (
	"net/http"

	"goji.io/pat"
)

func (rtr *Router) deletePlayer(w http.ResponseWriter, r *http.Request) {
	playerID := pat.Param(r, "player_id")

	err := rtr.Backend.DeletePlayer(r.Context(), playerID)
	if err != nil {
		rtr.handleErrorResponse(w, r, http.StatusInternalServerError, err)
		return
	}

	rtr.handleSuccessResponse(w, r, http.StatusNoContent, nil)
}
