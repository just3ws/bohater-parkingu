package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/just3ws/bohater-parkingu/models"
)

// // Index endpoint
// func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	fmt.Fprint(w, "Welcome!\n")
// }

// // LookupRate endpoint
// func LookupRate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
// }

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		file, err := ioutil.ReadFile("static/spec.json")

		if err != nil {
			panic(err)
		}

		fmt.Fprintf(w, string(file))
	})

	http.HandleFunc("/rate", func(w http.ResponseWriter, r *http.Request) {
		var starts time.Time
		var ends time.Time

		starts, _ = time.Parse(time.RFC3339, "2016-08-29T08:00:00.000Z")
		ends, _ = time.Parse(time.RFC3339, "2016-08-29T09:00:00.000Z")

		rate := models.Rate{Price: 1234, Starts: starts, Ends: ends}

		json.NewEncoder(w).Encode(rate)
	})

	log.Fatal(http.ListenAndServe(":8888", nil))
}
