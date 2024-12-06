package main

import (
	"fmt"
	"strings"
)

func main() {
	// TODO

	courses := map[int]string{

		1: "Calculus",
		2: "Biology",
		3: "Chemistry",
		4: "Computer Science",
		5: "Communications",
		6: "English",
		7: "Cantonese",
	}

	for id, courses := range courses {

		if strings.HasPrefix(courses, "C") {
			fmt.Println(id, courses)
		}
	}

	courses[4] = "Algorithms"
	courses[8] = "Spanish"
	delete(courses, 1)
	fmt.Print(courses)
}
