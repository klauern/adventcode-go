package main

import "testing"

func TestTwoChars(t *testing.T) {
	cnt := CountChars(`""`)
	if cnt.memChars != 2 || cnt.codeChars != 0 {
		t.Errorf("Not correct size, %+v", cnt)
	}
}


func TestAbcChars(t *testing.T) {
	//cnt := CountChars(`"abc"`)
}