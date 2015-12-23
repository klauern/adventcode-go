package main

import (
	"bufio"
	"os"
	"strings"
	"fmt"
	"strconv"
	"math"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var totalFt int = 0
	for scanner.Scan() {
		line := scanner.Text()
		space := strings.Split(line, "x")
		length, _ := strconv.Atoi(space[0])
		width, _ := strconv.Atoi(space[1])
		height, _ := strconv.Atoi(space[2])
		lenw := length * width
		widhei := width * height
		heilen := height * length
		total :=  2*lenw + 2*widhei + 2*heilen
		arr := []int {
			lenw,
			widhei,
			heilen,
		}
		// find smallest side
		minarea := math.MaxInt64
		for _, value := range arr {
			if  value < minarea {
				minarea = value
			}
		}
		totalFt += total
		totalFt += minarea
	}

	fmt.Println(totalFt)
}