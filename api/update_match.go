package api

import (
	"net/http"

	"github.com/benspotatoes/extrawatch/models"
)

func (rtr *Router) updateMatch(w http.ResponseWriter, r *http.Request) {
	err := rtr.Backend.UpdateMatch(r.Context(), &models.UpdateParams{})
	if err != nil {
		w.WriteHeader(http.StatusNotImplemented)
		return
	}

	w.WriteHeader(http.StatusOK)
}
