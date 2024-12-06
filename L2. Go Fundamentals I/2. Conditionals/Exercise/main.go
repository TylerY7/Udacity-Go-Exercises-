package main

import "fmt"

func main() {
	// TODO

	number := 500

	if number < 0 {
		fmt.Println(number, "is negative")
	} else if number < 100 {
		fmt.Println(number, "is positive")
	} else {
		fmt.Println(number, "is positive and is a large number!")
	}

}
