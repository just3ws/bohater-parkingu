package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func indexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	setCors(w)
	fmt.Fprintf(w, "MAYBE THIS IS THE RESTFUL API")
}

// used for COR preflight checks
func corsHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	setCors(w)
}

// util
func getFrontendURL() string {
	return "http://localhost:8080"
}

func setCors(w http.ResponseWriter) {
	frontendURL := getFrontendURL()
	w.Header().Set("Access-Control-Allow-Origin", frontendURL)
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

// Temporary Canary test to make sure Travis-CI is working
func Canary(word string) string {
	return word
}

func main() {
	log.SetOutput(os.Stdout)

	router := httprouter.New()
	router.GET("/", indexHandler)
	router.OPTIONS("/*any", corsHandler)

	log.Println("Running api server in dev mode")

	http.ListenAndServe(":8080", nil)
}

// NOOO!
