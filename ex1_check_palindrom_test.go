package main

import (
	"testing"
)

func TestPalindrome(t *testing.T) {
	cases := []struct {
		val  string
		want bool
	}{
		{val: "", want: true},
		{val: "a", want: true},
		{val: "hello world", want: false},
		{val: "1234", want: false},
		{val: "malayalam", want: true},
		{val: "Malayalam", want: true},
	}

	for _, i := range cases {
		if checkPalindrome(i.val) != i.want {
			resultString := "not a Palindrome"
			if i.want {
				resultString = "a Palindrome"
			}
			t.Errorf("Test %s, Expecting its %s :But Its failed", i.val, resultString)
		}
	}
}
