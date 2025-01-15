package main

import "fmt"

type Rectangle struct {
	length, width float64
}

type Square struct {
	side float64
}

type Perimeter interface {
	perimeter() float64
}

func (r Rectangle) perimeter() float64 {
	return (2 * r.length) + (2 * r.width)
}

func (s Square) perimeter() float64 {
	return s.side * 4
}
func getPerimeter(p Perimeter) float64 {
	return p.perimeter()
}

func main() {
	rectangle := Rectangle{length: 2, width: 4}
	square := Square{side: 2}

	fmt.Println("Rectangle perimeter: ", rectangle.perimeter())
	fmt.Println("Square perimeter: ", square.perimeter())
}
