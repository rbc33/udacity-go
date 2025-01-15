package main

import "fmt"

type Student struct {
	id   int
	name string
}
type Classroom struct {
	id          int
	capacity    int
	subject     string
	studentList []Student
}

func main() {
	dave := Student{1, "Dave"}
	cole := Student{2, "Cole"}
	c1 := Classroom{
		id:       1,
		capacity: 20,
		subject:  "Biology",
		studentList: []Student{
			dave,
			cole,
		},
	}
	c2 := new(Classroom)
	c2.id = 1
	c2.capacity = 20
	c2.subject = "Math"
	c2.studentList = []Student{
		dave,
		cole,
	}

	fmt.Printf("%v,%T\n", c1, c1)
	fmt.Printf("%v,%T\n", c2, c2)

}
