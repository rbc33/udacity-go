package main

import (
	"fmt"
	"net/http"
)

var cities = []string{"Tokyo", "Delhi", "Shanghai", "Sao Paulo", "Mexico City"}

type Handler interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hi there!!")
}

func cityList(w http.ResponseWriter, r *http.Request) {
	for _, city := range cities {
		fmt.Fprintf(w, "hi there!!%v\n", city)
	}
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/cityList", cityList)
	fmt.Println("Server is staarting on port 3000...")
	http.ListenAndServe("0.0.0.0:3000", nil)

}
