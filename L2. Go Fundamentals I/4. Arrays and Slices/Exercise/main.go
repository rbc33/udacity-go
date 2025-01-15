package main

import "fmt"

func main() {
	vegetableArr := [2]string{"Broccoli", "Carrot"}

	fmt.Println(vegetableArr)
	fmt.Println(vegetableArr[1])

	num := [5]int{1, 2, 3, 4, 5}
	fmt.Println(num[0:4])
}
