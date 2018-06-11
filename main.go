package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/just3ws/bohater-parkingu/models"
	"github.com/sirupsen/logrus"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	log.Debug("GET /")

	w.Header().Set("Content-Type", "application/json")

	file, err := ioutil.ReadFile("static/spec.json")

	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, string(file))
}

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

var log = logrus.New()

func main() {
	// log.SetOutput(os.Stdout)
	log.Out = os.Stdout
	log.Level = logrus.DebugLevel
	log.Debug("Server starting...")

	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/rate", RateLookupHandler)

	log.Fatal(http.ListenAndServe(":8888", nil))
}
