package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var wires map[string]string
//var totalLookupCount int = 0

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
		wires[strings.Trim(args[1], " ")] = strings.Trim(args[0], " ")
	}
	//for k, _ := range wires {
		//fmt.Printf("Derive %v\n", k)
		//fmt.Printf("%v: %v\n", k, deriveValue(k))
	//}
	fmt.Println(deriveValue("a"))
}

func deriveValue(key string) uint16 {
	fmt.Printf("DeriveValue(%v)\n", key)
	switch string(rune(key[0])) {
	case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
		val, _ := strconv.ParseUint(key, 10, 16)
		return uint16(val)
	default:
		return deriveInstruction(key)
	}
	panic("not found")
}

func deriveInstruction(key string) uint16 {
	fmt.Printf("deriveInstruction(%v)\n", key)
	pieces := strings.Split(wires[key], " ")
	switch len(pieces) {
	case 1:
		return deriveValue(pieces[0])
	case 2:
		if pieces[0] == "NOT" {
			return ^deriveValue(pieces[1])
		} else {
			panic(pieces)
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
}
