package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

var wires map[string]string

func init() {
	wires = make(map[string]string)
}

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		args := strings.Split(line, " -> ")
		// fill graph
		wires[args[1]] = args[0]
	}
	fmt.Println(deriveInstruction(wires["a"]))
}

func deriveValue(wire string) uint16 {
	val, err := strconv.ParseUint(wire, 10, 16)
	if err != nil {
		deriveInstruction(wire)
	} else {
		fmt.Printf("Value found: %v\n", val)
		return uint16(val)
	}
	fmt.Printf("Not found: %v\n", wire)
	panic("not found")
	return uint16(0)
}

func deriveInstruction(place string) uint16 {
	instruction := wires[place]
	pieces := strings.Split(instruction, " ")
	fmt.Printf("pieces split is %+v\n", pieces)
	switch len(pieces) {
	case 1:
		fmt.Printf("Deriving %v\n", pieces[0])
		return deriveValue(pieces[0])
	case 2:
		if pieces[0] == "NOT" {
			return ^deriveValue(pieces[1])
		}
	case 3:
		first := deriveValue(pieces[0])
		second := deriveValue(pieces[2])
		switch pieces[1] {
		case "LSHIFT":
			return first << second
		case "RSHIFT":
			return first >> second
		case "AND":
			return first & second
		case "OR":
			return first | second
		}
	}
	fmt.Printf("Not found: %v", place)
	panic("not found")
	return uint16(0)
}