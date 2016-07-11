package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var number_of_inputs int
	phonebook := make(map[string]int)
	var name string
	var phonenumber int
	fmt.Scan(&number_of_inputs)
	var i int

	for i = 0; i < number_of_inputs; i++ {
		// add things to phone book
		//read in a line
		// split on space
		fmt.Scan(&name)
		fmt.Scan(&phonenumber)
		phonebook[name] = phonenumber
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		// ok this is weird
		if _, ok := phonebook[line]; ok { // multiple assignment is weird. I guess every statement returns a bool if it succeeded or not
			fmt.Printf("%s=%v\n", line, phonebook[line])
		} else {
			fmt.Println("Not found")
		}
	}

}
