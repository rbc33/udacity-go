package main

import (
	"fmt"
	"strconv"
)

type Car struct {
	make string
	year uint16
	used bool
}

func (c Car) describe() string {
	used := ""
	if !c.used {
		used = " new "
	} else {
		used = " used  "
	}
	return "This " + strconv.Itoa(int(c.year)) + " " + c.make + " is a" + used + "car"
}

func main() {
	car1 := Car{make: "DeLorean", year: 1985, used: true}
	car2 := Car{make: "LaFerrari", year: 2023, used: false}

	fmt.Println(car1.describe())
	fmt.Println(car2.describe())

}
