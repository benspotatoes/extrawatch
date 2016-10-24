package api

import (
	"encoding/json"
	"net/http"

	"github.com/benspotatoes/extrawatch/models"
)

const (
	limitParam  = "limit"
	offsetParam = "offset"
)

func (rtr *Router) indexMatch(w http.ResponseWriter, r *http.Request) {
	// query := r.URL.Query()
	res, err := rtr.Backend.IndexMatch(r.Context(), &models.IndexParams{})
	if err != nil {
		w.WriteHeader(http.StatusNotImplemented)
		return
	}

	blob, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(blob)
}
