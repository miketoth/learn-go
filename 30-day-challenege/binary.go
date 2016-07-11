package main

import (
	"fmt"
)

func main() {
	var input int
	fmt.Scan(&input)
	var lsd int
	var current_count int
	var max int

	for input > 0 {
		lsd = input % 2
		input = input / 2
		if lsd == 1 {
			current_count++
			if current_count > max {
				max = current_count
			}
		} else {
			current_count = 0
		}
	}

	fmt.Println(max)
}
