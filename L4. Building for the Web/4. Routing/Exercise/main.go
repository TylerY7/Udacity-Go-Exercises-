package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var members = map[string]string{
	"1": "Andy",
	"2": "Peter",
	"3": "Gabriella",
	"4": "Jordy",
}

func getMembers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(members)
}

func deleteMember(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]

	if _, exists := members[id]; exists {
		delete(members, id)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(members)
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(members)
	}
}

func main() {
	// TODO
	router := mux.NewRouter()

	router.HandleFunc("/members", getMembers).Methods("GET")
	router.HandleFunc("/deleteMember/{id}", deleteMember).Methods("DELETE")

	fmt.Println("Server is starting on port 3000...")
	http.ListenAndServe(":3000", router)
}
