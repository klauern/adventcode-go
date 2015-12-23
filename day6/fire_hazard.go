package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

var lights [1000][1000]bool

//func init() {
//	lights = make([][]bool, 1000)
//	for i := 0; i < 1000; i++ {
//		lights[i] = make([]bool, 1000)
//	}
//}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "turn off") {
			rng := strings.TrimLeft(line, "turn off ")
			lightSlice := getRange(rng)
			turnOffLights(lightSlice)
		} else if strings.HasPrefix(line, "turn on") {
			rng := strings.TrimLeft(line, "turn on ")
			lightSlice := getRange(rng)
			turnOnLights(lightSlice)
		} else if strings.HasPrefix(line, "toggle") {
			rng := strings.TrimLeft(line, "toggle ")
			lightSlice := getRange(rng)
			toggleLights(lightSlice)
		}
	}
	lit := 0
	for i, _ := range lights {
		for _, isLit := range lights[i] {
			if isLit {
				lit += 1
			}
		}
	}
	fmt.Println(lit)
}

type lightRange struct {
	startx, starty int
	endx, endy     int
}

func getRange(through string) *lightRange {
	strs := strings.Split(through, " through ")
	starts := strings.Split(strs[0], ",")
	ends := strings.Split(strs[1], ",")
	startx, _ := strconv.Atoi(starts[0])
	starty, _ := strconv.Atoi(starts[1])
	endx, _ := strconv.Atoi(ends[0])
	endy, _ := strconv.Atoi(ends[1])
	// Totally forgot that people could put in ranges that walk backwards, like 999,0 777,0, meaning everything between 777 and 999 get filled or something
	lrange := &lightRange{
		startx: startx,
		starty: starty,
		endx:   endx,
		endy:   endy,
	}
	fmt.Println(lrange)
	return lrange
}

func turnOffLights(lightSlice *lightRange) {
	for i := lightSlice.startx; i < lightSlice.endx; i++ {
		for j := lightSlice.starty; j < lightSlice.endy; j++ {
			lights[i][j] = false
		}
	}
}

func turnOnLights(lightSlice *lightRange) {
	for i := lightSlice.startx; i < lightSlice.endx; i++ {
		for j := lightSlice.starty; j < lightSlice.endy; j++ {
			lights[i][j] = true
		}
	}
}

func toggleLights(lightSlice *lightRange) {
	for i := lightSlice.startx; i < lightSlice.endx; i++ {
		for j := lightSlice.starty; j < lightSlice.endy; j++ {
			lights[i][j] = !lights[i][j]
		}
	}
}