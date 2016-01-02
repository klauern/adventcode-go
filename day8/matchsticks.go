package main

import 	(
	"strconv"
	"strings"
	"os"
	"bufio"
	"fmt"
)

func main() {
	current := &CharCount{}


	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		lineCnt := CountChars(line)
		current.memChars += lineCnt.memChars
		current.codeChars += lineCnt.codeChars
		fmt.Printf("Current Counts are %+v\n", current)
		//fmt.Printf("Count of %s is %+v", line, CountChars(line))
	}

	fmt.Printf("Total %+v\n", current)
	fmt.Printf("Difference is %v", current.codeChars-current.memChars)
}

type CharCount struct {
	codeChars int
	memChars  int
}

func CountChars(line string) *CharCount {
	unquoted, err := strconv.Unquote(line)
	if err != nil {
		panic(err)
	}

	unquotedReader := strings.NewReader(unquoted)

	return &CharCount{
		codeChars: len(line),
		memChars:  unquotedReader.Len(),
	}
}
