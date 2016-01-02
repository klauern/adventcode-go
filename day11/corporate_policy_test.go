package main

import "testing"

var incrementTests = []struct {
	in  string
	out string
}{
	{"a", "b"},
	{"az", "ba"},
	{"abc", "abd"},
	{"abz", "aca"},
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
	{"aabb", true},
	{"aaab", false},
	{"abccdeff", true},
	{"abcdefg", false},
}

func TestContainsLetterPair(t *testing.T) {
	for _, tt := range letterPairs {
		pass := ContainsTwoLetterPairs(tt.in)
		if pass != tt.out {
			t.Errorf("ContainsLetterPair(\"%s\"): \tWanted %6s, Got %-6s", tt.in, tt.out, pass)
		}
	}
}

var invalidStrings = []struct {
	in  string
	out bool
}{
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

var increasingLetters = []struct {
	in  string
	out bool
}{
	{"abc", true},
	{"febecde", true},
	{"hijklmmn", true},
	{"abbcegjk", false},
	{"abbceffg", false},
}

func TestIncreasingLetters(t *testing.T) {
	for _, tt := range increasingLetters {
		pass := HasIncreasingLetters(tt.in)
		if pass != tt.out {
			t.Errorf("HasIncreasingLetters(\"%s\"): \tWanted %v, Got %v", tt.in, tt.out, pass)
		}
	}
}

var nextPasswords = []struct {
	in  string
	out string
}{
	{"abcdefgh", "abcdffaa"},
	{"ghijklmn", "ghjaabcc"},
}

func TestNextPassword(t *testing.T) {
	for _, tt := range nextPasswords {
		pass := NextPassword(tt.in)
		if pass != tt.out {
			t.Errorf("NextPassword(\"%s\"): \tWanted %s, Got %s", tt.in, tt.out, pass)
		}
	}
}
