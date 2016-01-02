package main

import (
	"testing"
	"os"
"bufio"
	"fmt"
)

func TestFile(t *testing.T) {

		file, err := os.Open("test_input.txt")
		if err != nil {
			panic(err)
		}
		scanner := bufio.NewScanner(file)
		//scanner.Split(bufio.ScanLines)
			t.Logf("Shortest Route is is %+v", FindShortest(scanner))
			fmt.Printf("Shortest Route is %+v", FindShortest(scanner))
}
