package handlers

import (
	"encoding/json"
	"net/http"

	"just3ws/bohater-parkingu/models"
	"just3ws/bohater-parkingu/params"

	log "github.com/sirupsen/logrus"
)

// RateHandler for "/rate"
func RateHandler(w http.ResponseWriter, r *http.Request) {
	log.Debug("GET /rate")

	w.Header().Set("Content-Type", "application/json")

	starts, err := params.StartsParam(r.URL)
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ends, err := params.EndsParam(r.URL)
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var price int32 = 8888
	rate := models.Rate{Price: price, Starts: starts, Ends: ends}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(rate)
}
