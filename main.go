package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/just3ws/bohater-parkingu/handlers"
	"github.com/just3ws/bohater-parkingu/models"
	log "github.com/sirupsen/logrus"
)

func EndsParam(url *url.URL) (time.Time, error) {
	params, ok := url.Query()["ends"]

	if !ok || len(params) < 1 {
		msg := "ends: param is missing"
		return time.Time{}, errors.New(msg)
	}

	ends, err := time.Parse(time.RFC3339, params[0])
	if err != nil {
		msg := "ends: param is invalid"
		return time.Time{}, errors.New(msg)
	}

	return ends, nil
}

func StartsParam(url *url.URL) (time.Time, error) {
	params, ok := url.Query()["starts"]

	if !ok || len(params) < 1 {
		msg := "starts: param is missing"
		return time.Time{}, errors.New(msg)
	}

	starts, err := time.Parse(time.RFC3339, params[0])
	if err != nil {
		msg := "starts: param is invalid"
		return time.Time{}, errors.New(msg)
	}

	return starts, nil
}

func RateLookupHandler(w http.ResponseWriter, r *http.Request) {
	log.Debug("GET /rate")

	w.Header().Set("Content-Type", "application/json")

	starts, err := StartsParam(r.URL)
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ends, err := EndsParam(r.URL)
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
	http.HandleFunc("/rate", RateLookupHandler)

	log.Fatal(http.ListenAndServe(":8888", nil))
}
