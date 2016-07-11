package main

import (
	"fmt"
)

func main() {
	var input int
	fmt.Scan(&input)
	output := factorial(input)
	fmt.Println(output)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
