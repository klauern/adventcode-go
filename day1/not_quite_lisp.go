package main

import (
	"io/ioutil"
	"fmt"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	var floor int = 0
	for _, c := range string(input) {
		if c == '(' {
			floor += 1
		}
		if c == ')' {
			floor -= 1
		}
	}
	fmt.Printf("floor: %v", floor)
}
