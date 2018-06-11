package main

import (
	"net/http"
	"os"

	"github.com/just3ws/bohater-parkingu/handlers"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func main() {
	log.Debug("Server starting...")

	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/rate", handlers.RateHandler)

	http.ListenAndServe(":8888", nil)
}
