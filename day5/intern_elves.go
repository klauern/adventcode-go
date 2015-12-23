package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var BannedStrings = []string{"ab", "cd", "pq", "xy"}
var Vowels = []string{"a", "e", "i", "o", "u"}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	niceGuys := 0
	for scanner.Scan() {
		line := scanner.Text()
		if threeVowels(line) && twiceInARow(line) && !notContains(line) {
			niceGuys += 1
		}
	}
	fmt.Println(niceGuys)
}

func threeVowels(input string) bool {
	count := 0
	for _, val := range Vowels {
		count += strings.Count(input, val)
	}
	if count >= 3 {
		return true
	}
	return false
}

func twiceInARow(input string) bool {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanRunes)
	prevChar := ""
	for scanner.Scan() {
		curChar := scanner.Text()
		if curChar == prevChar {
			return true
		}
		prevChar = curChar
	}
	return false
}

func notContains(input string) bool {
	for _, val := range BannedStrings {
		if strings.Contains(input, val) {
			return true
		}
	}
	return false
}
