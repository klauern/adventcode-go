package main

import "testing"

var equalityTests = []struct {
	in  string
	out int
}{
	{`[1,2,3]`, 6},
	{`{"a":2,"b":4}`, 6},
	{`[[[3]]]`, 3},
	{`{"a":{"b":4},"c":-1}`, 3},
	{`{"a":[-1,1]}`, 0},
	{`[-1,{"a":1}]`, 0},
	{`[]`, 0},
	{`{}`, 0},
}

func TestEquality(t *testing.T) {
	for _, tt := range equalityTests {
		sum := SumString(tt.in)
		if sum != tt.out {
			t.Errorf("SumString(\"%s\"): \tWanted %6s, Got %-6s", tt.in, tt.out, sum)
		}
	}
}
