package main

import (
	"strconv"
	"fmt"
)

func main() {
	input := "1321131112"
	for i := 0; i < 40; i++ {
		newInput := LookAndSay(input)
		input = newInput
		fmt.Printf("Iteration %2d: output %d\n", i, len(input))
	}

}

func LookAndSay(input string) string {
	counting := ""
	output := ""
	size := 0
	for _, r := range input {
		c := string(r)
		if counting == "" {
			counting = c
		}
		if c != counting {
			append := fmt.Sprintf("%s%s", strconv.Itoa(size), counting)
			output += append
			size = 0
			counting = c
		}
		size += 1
	}
	append := fmt.Sprintf("%s%s", strconv.Itoa(size), counting)
	output += append
	return output
}
