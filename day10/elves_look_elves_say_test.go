package main

import "testing"

var sequenceTests = []struct {
	in string
	out string
}{
	{"1", "11"},
	{"11", "21"},
	{"21", "1211"},
	{"1211", "111221"},
	{"111221", "312211"},
}

func TestLookAndSay(t *testing.T) {
	for _, tt := range sequenceTests {
		said := LookAndSay(tt.in)
		if said != tt.out {
			t.Errorf("LookAndSay(\"%s\"): \tWanted %6s, Got %-6s", tt.in, tt.out, said)
		}
	}
}
