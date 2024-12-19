package main

import (
	"fmt"
	"net/http"
)

var cities = []string{"Tokyo", "Delhi", "Shanghai", "Sao Paulo", "Mexico City"}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello welcome :3")

}

func cityList(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "List of most populous cities")

	for _, city := range cities {
		fmt.Fprintf(w, "%s\n", city)
	}
}

func main() {
	// TODO

	http.HandleFunc("/", index)
	http.HandleFunc("/cityList", cityList)

	http.ListenAndServe(":3000", nil)

}
