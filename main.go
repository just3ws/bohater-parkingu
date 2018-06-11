package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/just3ws/bohater-parkingu/handlers"
	"github.com/just3ws/bohater-parkingu/models"
	"github.com/just3ws/bohater-parkingu/params"
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

	rate := models.Rate{Price: 1234, Starts: starts, Ends: ends}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(rate)
}

// var log = logrus.New()

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.TextFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}

func main() {
	// log.Out = os.Stdout
	// log.Level = logrus.DebugLevel

	log.Debug("Server starting...")

	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/rate", RateHandler)

	log.Fatal(http.ListenAndServe(":8888", nil))
}
