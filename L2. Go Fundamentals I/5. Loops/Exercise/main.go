package main

import "fmt"

func fizzbuzz(n int) []string {
	sliceStrings := []string{}

	for i := 1; i <= n; i++ {

		if i%3 == 0 && i%5 == 0 {
			sliceStrings = append(sliceStrings, "Fizzbuzz")
		} else if i%5 == 0 {
			sliceStrings = append(sliceStrings, "Buzz")
		} else if i%3 == 0 {
			sliceStrings = append(sliceStrings, "Fizz")
		} else {
			stringValue := fmt.Sprintf("%d", i)
			sliceStrings = append(sliceStrings, stringValue)
		}

	}

	return sliceStrings
}

func main() {
	// TODO

	fmt.Println(fizzbuzz(30))

}
