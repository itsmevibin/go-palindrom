package main

import (
	"fmt"
	"strings"
)

func reverseString(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func checkPalindrome(s string) (result int) {
	result = strings.Compare(s, reverseString(s))
	return
}

func main() {
	fmt.Println(reverseString("hello world"))
}
