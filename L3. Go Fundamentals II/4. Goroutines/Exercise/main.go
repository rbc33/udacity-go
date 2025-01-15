package main

import (
	"fmt"
	"time"
)

func iterPrint(s []string) {
	for _, str := range s {
		fmt.Println(str)
	}

}

func main() {
	colorNames := []string{"red", "orange", "yellow", "green", "blue", "indigo", "violet"}
	colorCodes := []string{"#FF0000", "#FF7F00", "#FFFF00", "#00FF00", "#0000FF", "#4B0082", "#9400D3"}
	start := time.Now()
	// iterPrint(colorNames)
	// iterPrint(colorCodes)
	go iterPrint(colorNames)
	go iterPrint(colorCodes)
	duration := time.Since(start)
	fmt.Println(duration)
}
