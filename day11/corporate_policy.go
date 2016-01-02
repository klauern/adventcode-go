package main

import "strings"

func main() {

}

func NextPassword(current string) string {
	return ""
}

func HasIncreasingLetters(password string) bool {
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

func ContainsLetterPair(password string) bool {
	previous := ""
	for _, r := range password {
		current := string(r)
		if current == previous {
			return true
		}
		previous = current
	}
	return false
}

func Increment(password string) string {
	out := []rune(password)
	for i := len(out)-1; i >= 0; i-- {
		if out[i] == 'z' {
			out[i] = 'a'
		} else {
			out[i] = out[i]+1
			break
		}
	}
	return string(out)
}
