package main

import (
	"bytes"
	"fmt"
)

func main() {
	var number_of_inputs int
	var temp_string string
	var count int
	fmt.Scan(&number_of_inputs)
	count = 0
	i := number_of_inputs
	j := 0
	input_strings := make([]string, number_of_inputs)

	// first get all inputs
	for i > 0 {
		fmt.Scan(&temp_string)
		input_strings[count] = temp_string
		i--
		count++
	}
	fmt.Println(input_strings)

	// then print output
	for i = 0; i < number_of_inputs; i++ {
		var even_buffer bytes.Buffer
		var odd_buffer bytes.Buffer
		for j = 0; j < len(input_strings[i]); j++ {
			if j%2 == 0 {
				// even
				even_buffer.WriteString(string(input_strings[i][j]))
			} else {
				// odd
				odd_buffer.WriteString(string(input_strings[i][j]))
			}
		}
		// print out even buffer SPACE odd buffer
		fmt.Printf("%s %s\n", even_buffer.String(), odd_buffer.String())
	}
}
