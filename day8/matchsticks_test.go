package main

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

var countTests = []struct {
	in  string
	out *CharCount
}{
	{`""`, &CharCount{codeChars: 2, memChars: 0}},
	{`"abc"`, &CharCount{codeChars: 5, memChars: 3}},
	{`"aaa\"aaa"`, &CharCount{codeChars: 10, memChars: 7}},
	{`"\x27"`, &CharCount{codeChars: 6, memChars: 1}},
}

func TestCharCounting(t *testing.T) {
	for _, tt := range countTests {
		cnt := CountChars(tt.in)
		if cnt.memChars != tt.out.memChars || cnt.codeChars != tt.out.codeChars {
			t.Errorf("CodeChars(%s) => %+v, want %+v", tt.in, cnt, tt.out)
		}
	}
}

func TestCharCountingFile(t *testing.T) {
	file, err := os.Open("test_input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		t.Logf("Count of %s is %+v", line, CountChars(line))
		fmt.Printf("Count of %s is %+v", line, CountChars(line))
	}
}
