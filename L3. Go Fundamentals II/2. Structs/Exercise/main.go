package main

import "fmt"

type student struct {
	id      int
	student string
}

type classroom struct {
	id          int
	capacity    int
	subject     string
	studentList []student
}

func main() {
	// TODO

	c1 := classroom{
		id:       01,
		capacity: 20,
		subject:  "Computer Science",
		studentList: []student{
			{
				id:      01,
				student: "Tyler",
			},
			{
				id:      02,
				student: "Josh",
			},
		},
	}
	fmt.Println(c1)

	c2 := new(classroom)
	c2.id = 02
	c2.capacity = 30
	c2.subject = "Art"
	c2.studentList = []student{
		{
			id:      01,
			student: "Georgia",
		},
		{
			id:      02,
			student: "Winters",
		},
	}

	fmt.Println(c2)

}
