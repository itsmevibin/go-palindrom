package main

import (
	"strings"
	"testing"
)

func TestReverse(t *testing.T) {
	s := "Hello word"
	revR := reverseString(s)

	if strings.Compare(revR, "drow olleH") != 0 {
		t.Errorf("Expecting: %s\n. But got: %s\n", s, revR)
	}
}

func TestPalindrome(t *testing.T) {
	if checkPalindrome("") != 0 {
		t.Errorf("Reverse string is failed")
	}
}
