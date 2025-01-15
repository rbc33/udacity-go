package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var dictionary = map[string]string{
	"Go":     "A programing language created by Google.",
	"Gopher": "A software engenier who build with Go.",
	"Golang": "Another name for Go",
}

//	If our handler function is meant to return HTML, we can actually use the same syntax to do so!
//
// w.Header().Set("Content-Type", "text/html; charset=utf-8")
func getDictionary(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 201
	json.NewEncoder(w).Encode(dictionary)
}

func main() {
	http.HandleFunc("/", getDictionary)
	fmt.Println("Server is staarting on port 3000...")
	http.ListenAndServe("0.0.0.0:3000", nil)
}
