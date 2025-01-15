package main

import (
	"fmt"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	fmt.Println("Serving is starting on port 3000...")
	http.ListenAndServe("0.0.0.0:3000", nil)
}
