package api

import (
	"net/http"

	"github.com/benspotatoes/extrawatch/models"
)

func (rtr *Router) postMatch(w http.ResponseWriter, r *http.Request) {
	err := rtr.Backend.InsertMatch(r.Context(), &models.InsertParams{})
	if err != nil {
		w.WriteHeader(http.StatusNotImplemented)
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
