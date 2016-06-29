package main

import (
	"fmt"
	"os"
)

func main() {
	var count int = 0
	var size int
	var temp int

	fmt.Scan(&size)
	number_array := make([]int, size)

	for count < size {
		fmt.Scan(&temp)
		number_array[count] = temp

		if count == size-1 {
			// reverse
			for count >= 0 {
				fmt.Print(number_array[count])
				if count > 0 {
					fmt.Print(" ")
				}
				count--
			}
			os.Exit(0)
		}
		count++
	}
}
