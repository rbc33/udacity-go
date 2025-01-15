package main

import (
	"fmt"
	"strings"
)

func main() {
	courses := map[int]string{1: "Calculus", 2: "Biology", 3: "Chemistry", 4: "Computer Science", 5: "Communications", 6: "English", 7: "Cantonese"}
	for _, class := range courses {
		if strings.HasPrefix(class, "C") {
			fmt.Println(class)
		}
	}
	courses[4] = "Algorithms"
	courses[8] = "Spanish"
	delete(courses, 1)
	fmt.Println(courses)
}
