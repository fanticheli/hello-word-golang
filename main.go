package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleHelloName(w http.ResponseWriter, rq *http.Request) {
	vars := mux.Vars(rq)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": fmt.Sprintf("Hello %s", vars["name"]),
	})
}

func HandleHello(w http.ResponseWriter, rq *http.Request) {
	fmt.Fprintf(w, "Hello world!\n")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HandleHello)
	r.HandleFunc("/{name}", HandleHelloName).Methods("GET")
	http.ListenAndServe(":8080", r)
}
