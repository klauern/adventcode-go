package main

import (
	"strings"
	"fmt"
)

func main() {
	fmt.Println("Next Password for \"hepxcrrq\"", NextPassword("hepxcrrq"))
}

func NextPassword(current string) string {
	password := current
	for  {
		password = Increment(password)
		if HasIncreasingLetters(password) && !ContainsInvalidChars(password) && ContainsTwoLetterPairs(password) {
			return password
		}
	}
	return ""
}

func HasIncreasingLetters(password string) bool {
	out := []rune(password)
	for i := 0; i <= len(out)-3; i++ {
		// && out[i+2]+1 == out[i+3]
		if out[i]+1 == out[i+1] && out[i+1]+1 == out[i+2] {
			return true
		}
	}
	return false
}

func ContainsInvalidChars(password string) bool {
	for _, r := range "iol" {
		c := string(r)
		if strings.Contains(password, c) {
			return true
		}
	}
	return false
}

func ContainsTwoLetterPairs(password string) bool {
	previous := ""
	pairs := 0
	matched := ""
	for _, r := range password {
		current := string(r)
		if current == previous && matched != current {
			pairs += 1
			matched = current
		}
		previous = current
	}
	return pairs >= 2
}

func Increment(password string) string {
	out := []rune(password)
	for i := len(out) - 1; i >= 0; i-- {
		if out[i] == 'z' {
			out[i] = 'a'
		} else {
			out[i] = out[i] + 1
			break
		}
	}
	return string(out)
}
