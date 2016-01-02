package main

import "strings"

func main() {


}


func NextPassword(current string) string {

}

func HasIncreasingLetters(password string) bool {

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

}