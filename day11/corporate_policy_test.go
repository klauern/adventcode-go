package main

import "testing"

func TestNextPassword(t *testing.T) {

}

var incrementTests = []struct {
	in  string
	out string
}{
	{"a", "b"},
	{"az", "ba"},
	{"abc", "abd"},
	{"abz", "aca"},
}

var nextPasswordTests = []struct {
	in  string
	out string
}{
	{"abcdefgh", "abcdffaa"},
	{"ghijklmn", "ghjaabcc"},
}

func TestIncrement(t *testing.T) {
	for _, tt := range incrementTests {
		pass := Increment(tt.in)
		if pass != tt.out {
			t.Errorf("Increment(\"%s\"): \tWanted %6s, Got %-6s", tt.in, tt.out, pass)
		}
	}
}

var letterPairs = []struct {
	in  string
	out bool
}{
	{"aa", true},
	{"abccdeff", true},
	{"abcdefg", false},
}

func TestContainsLetterPair(t *testing.T) {
	for _, tt := range letterPairs {
		pass := ContainsLetterPair(tt.in)
		if pass != tt.out {
			t.Errorf("ContainsLetterPair(\"%s\"): \tWanted %6s, Got %-6s", tt.in, tt.out, pass)
		}
	}
}

var invalidStrings = []struct {
in string
out bool
} {
	{"abcdefghi", true},
	{"hijklmmn", true},
	{"abbceffg", false},
}

func TestInvalidStrings(t *testing.T) {
	for _, tt := range invalidStrings {
		pass := ContainsInvalidChars(tt.in)
		if pass != tt.out {
			t.Errorf("ContainsInvalidChars(\"%s\"): \tWanted %6s, Got %-6s", tt.in, tt.out, pass)
		}
	}
}