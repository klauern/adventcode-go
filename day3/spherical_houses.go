package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	var x, y int = 0, 0
	houses := make(map[string]bool)
	houses[fmt.Sprintf("%d,%d", x, y)] = true
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		switch scanner.Text() {
		case "<":
			x -= 1
			fmt.Println("Left")
		case ">":
			x += 1
			fmt.Println("Right")
		case "^":
			y += 1
			fmt.Println("Up")
		case "v":
			y -= 1
			fmt.Println("Down")
		}
		houses[fmt.Sprintf("%d,%d", x, y)] = true
	}

	fmt.Printf("Houses hit: %v", len(houses))
}
