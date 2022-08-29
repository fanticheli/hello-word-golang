package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

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
	port := os.Getenv("PORT")

	if port == "" {
		port = viper.GetString("server.port")
	}

	r := mux.NewRouter()
	r.HandleFunc("/", HandleHello)
	r.HandleFunc("/{name}", HandleHelloName).Methods("GET")

	log.Printf("LISTEN ON PORT: %v", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), r)
}
