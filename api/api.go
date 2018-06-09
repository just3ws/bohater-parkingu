package main

import (
	"fmt"
	"net/http"
)

func main() {
	port := ":8888"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello there, YOU'VE REQUESTED: %s\n", r.URL.Path)
	})

	http.ListenAndServe(port, nil)
}
