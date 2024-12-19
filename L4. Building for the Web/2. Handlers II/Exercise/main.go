package main

import (
	"fmt"
	"net/http"
)

var cityPopulations = map[string]uint32{
	"Tokyo":       37435191,
	"Delhi":       29399141,
	"Shanghai":    26317104,
	"Sao Paulo":   21846507,
	"Mexico City": 21671908,
}

func index(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	for key, cityPop := range cityPopulations {
		fmt.Fprintf(w, "<h2>City %s: City's Pop: %d</h2>", key, cityPop)
	}

}

func main() {
	// TODO
	http.HandleFunc("/", index)

	fmt.Println("Server is starting on port 3000...")
	http.ListenAndServe(":3000", nil)
}
