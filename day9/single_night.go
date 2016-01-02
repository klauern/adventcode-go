package main

import (
"io"
"bufio"
)

type FlightDistances map[string][]string

type FlightPaths struct {
	flights FlightDistances
	distance
}

func main() {}



func FindShortest(r io.Reader) string {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {

	}


}