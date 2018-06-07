package main

import (
	"fmt"
	"net/http"
)

// Say hello
func Say() string {
	return "Hello, World!"
}

func main() {

	port := ":8888"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	fs := http.FileServer(http.Dir("web/static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(port, nil)
}
