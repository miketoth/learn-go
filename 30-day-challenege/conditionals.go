package main

import (
	"fmt"
)

func main() {
	var input int
	fmt.Scan(&input)

	if input >= 1 && input <= 100 {

		if input%2 != 0 {
			fmt.Println("Weird")
		} else {
			if input > 20 {
				fmt.Println("Not Weird")
			} else if input >= 2 && input <= 5 {
				fmt.Println("Not Weird")
			} else if input >= 6 && input <= 20 {
				fmt.Println("Weird")
			}
		}
	}
}
