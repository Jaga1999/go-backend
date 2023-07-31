package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

type ResponseData struct {
	Message string `json:"message"`
}

func getMethod(w http.ResponseWriter, r *http.Request) {
	response := ResponseData{
		Message: "Hi there, I love Golang!",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/get", getMethod)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
