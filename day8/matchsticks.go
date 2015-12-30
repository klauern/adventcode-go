package main

import "strings"

func main() {

}

type CharCount struct {
	codeChars int
	memChars  int64
}

func CharCountReplacer() *strings.Replacer {
	return strings.NewReplacer(`\"`, `"`, ``)
}

func CountChars(line string) *CharCount {
	lineReader := strings.NewReader(line)
	return &CharCount{
		codeChars: lineReader.Len(),
		memChars:  lineReader.Size(),
	}
}
