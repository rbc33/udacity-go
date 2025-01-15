package main

import (
	"fmt"
)

func getRectangleArea(b, h int) int {
	return b * h
}

func main() {
	area := getRectangleArea(5, 10)
	if area < 50 {
		fmt.Printf("The area is %v, which less than 50", area)
	} else {
		fmt.Printf("The area is %v, which grater than or equal to 50 \n", area)
	}

}
