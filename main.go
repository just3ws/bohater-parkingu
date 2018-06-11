package bohaterparkingu

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Index endpoint
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

// LookupRate endpoint
func LookupRate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func main() {
	http.HandleFunc("/v1/rate", func(w http.ResponseWriter, r *http.Request) {
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
