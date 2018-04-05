package main

import (
	"fmt"
	"strings"
)

func checkPalindrome(s string) bool {
	s = strings.ToLower(s)
	mid := len(s) / 2
	last := len(s) - 1
	for i := 0; i < mid; i++ {
		if s[i] != s[last-i] {
			return false
		}
	}
	return true
}

func main() {
	var inputVal string
	fmt.Println("Enter string")
	fmt.Scanf("%s\n", &inputVal)
	if checkPalindrome(inputVal) {
		fmt.Println("YES... Entered string is a palindrome")
	} else {
		fmt.Println("No..! Entered string is not a palindrome")
	}
}
