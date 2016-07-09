package main

import (
	"fmt"
)

func main() {
	var number_of_inputs int
	var input_string string
	fmt.Scan(&number_of_inputs)

	// first get all inputs
	for number_of_inputs > 0 {
		fmt.Scan(&input_string)
		fmt.Println(input_string)
		number_of_inputs--
	}

	// then print output
}
