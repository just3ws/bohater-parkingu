package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// IndexHandler for "/"
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	log.Debug("GET /")

	w.Header().Set("Content-Type", "application/json")

	file, err := ioutil.ReadFile("static/spec.json")

	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, string(file))
}
