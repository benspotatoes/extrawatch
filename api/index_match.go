package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/benspotatoes/extrawatch/models"
)

type indexMatchResp struct {
	Matches []*models.Match `json:"matches"`
}

const (
	limitParam  = "limit"
	offsetParam = "offset"
	filterParam = "filter"

	maxLimit = 30
)

var (
	errMaxLimit = errors.New("requested limit exceeded max")
)

func (rtr *Router) indexMatch(w http.ResponseWriter, r *http.Request) {
	var limit int
	var offset int
	var err error

	query := r.URL.Query()
	filter := query.Get(filterParam)

	limitp := query.Get(limitParam)
	if limitp != "" {
		limit, err = strconv.Atoi(limitp)
		if err != nil {
			rtr.handleErrorResponse(w, r, http.StatusBadRequest, err)
			return
		}
		if limit > maxLimit {
			rtr.handleErrorResponse(w, r, http.StatusBadRequest, errMaxLimit)
		}
	}
	offsetp := query.Get(offsetParam)
	if offsetp != "" {
		offset, err = strconv.Atoi(offsetp)
		if err != nil {
			rtr.handleErrorResponse(w, r, http.StatusBadRequest, err)
			return
		}
	}

	matches, err := rtr.Backend.IndexMatch(r.Context(), limit, offset, filter)
	if err == sql.ErrNoRows {
		rtr.handleErrorResponse(w, r, http.StatusNotFound, nil)
		return
	}
	if err != nil {
		rtr.handleErrorResponse(w, r, http.StatusInternalServerError, err)
		return
	}

	res := &indexMatchResp{Matches: matches}
	blob, err := json.Marshal(res)
	if err != nil {
		rtr.handleErrorResponse(w, r, http.StatusInternalServerError, err)
		return
	}

	rtr.handleSuccessResponse(w, r, http.StatusOK, blob)
}
