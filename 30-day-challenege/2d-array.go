package main

import (
	"fmt"
)

func main() {
	// given that input is 6x6
	var multi [6][6]int
	var i int
	var j int
	var temp int
	var max int

	for i = 0; i < 6; i++ {
		for j = 0; j < 6; j++ {
			fmt.Scan(&temp)
			multi[i][j] = temp
		}
	}

	for i = 0; i < 4; i++ {
		for j = 0; j < 4; j++ {
			// for (0,0) print the sum of (0,0)(0,1)(0,2)(1,1)(2,0)(2,1)(2,2)
			temp = multi[i][j] + multi[i][j+1] + multi[i][j+2] + multi[i+1][j+1] + multi[i+2][j] + multi[i+2][j+1] + multi[i+2][j+2]
			if i == 0 && j == 0 {
				max = temp
			} else if temp > max {
				max = temp
			}
		}
	}

	fmt.Println(max)
}
