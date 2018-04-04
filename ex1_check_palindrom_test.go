package main

import (
	"strings"
	"testing"
)

func TestReverse(t *testing.T) {
	cases := []struct {
		val  string
		want string
	}{
		{val: "", want: ""},
		{val: "a", want: "a"},
		{val: "hello world", want: "dlrow olleh"},
		{val: "1234", want: "4321"},
	}

	for _, i := range cases {
		if strings.Compare(reverseString(i.val), i.want) != 0 {
			t.Errorf("Expecting: %s\n. But got: %s\n", i.want, reverseString(i.val))
		}
	}
}

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
		{val: "Malayalamx", want: true},
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
