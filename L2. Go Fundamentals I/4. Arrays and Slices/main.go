package main

import "fmt"

func main() {
	// 0 1 1 2 3 5 8 13
	// var fibonacciArr [8]int

	// fibonacciArr[0] = 0
	// fibonacciArr[1] = 1
	// fibonacciArr[2] = 1
	// fibonacciArr[3] = 2
	// fibonacciArr[4] = 3
	// fibonacciArr[5] = 5
	// fibonacciArr[6] = 8
	// fibonacciArr[7] = 13

	// fibonacciArr = [8]int{0, 1, 1, 2, 3, 5, 8, 13}

	fibonacciArr := [8]int{0, 1, 1, 2, 3, 5, 8, 13}
	fmt.Println(fibonacciArr)
	fmt.Println(fibonacciArr[0])
	fmt.Println(fibonacciArr[7])
	fmt.Println(len(fibonacciArr))
	fmt.Println(fibonacciArr[0:4])
}
