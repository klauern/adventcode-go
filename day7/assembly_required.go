package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	fmt.Println(deriveValue(wires["a"]))
}

func deriveValue(wire string) uint16 {
	fmt.Printf("deriveValue(%v)\n", wire)
	switch string(wire[0]) {
	case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
		fmt.Println("Is Num")
		val, _ := strconv.ParseUint(wire, 10, 16)
		fmt.Printf("Value found: %v\n", val)
		return uint16(val)
	default:
		return deriveInstruction(wire)
	}
	fmt.Printf("Not found: %v\n", wire)
	panic("not found")
	return uint16(0)
}

func deriveInstruction(place string) uint16 {
	fmt.Printf("deriveInstruction(%v)\n", place)
	fmt.Printf("Deriving Instruction %v\n", place)
	instruction := wires[place]
	fmt.Printf("Instruciton %v lookup is %v\n", place, instruction)
	pieces := strings.Split(instruction, " ")
	fmt.Printf("pieces split is %+v\n", pieces)
	//value := uint16(0)
	switch len(pieces) {
	case 1:
		fmt.Printf("Length 1: %v\n", pieces[0])
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

	panic("Not supposed to fall through")
	return uint16(0)
}
