package main

import "strings"

func main() {

}

type CharCount struct {
	codeChars int64
	memChars  int
}

func CharCountReplacer() *strings.Replacer {
	return strings.NewReplacer(`\"`, `"`, ``)
}

func CountChars(line string) *CharCount {
	lineReader := strings.NewReader(line)
	return &CharCount{
		codeChars: lineReader.Size(),
		memChars:  lineReader.Len(),
	}
}
