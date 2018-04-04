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
	var inputVal string
	fmt.Println("Enter string:")
	fmt.Scanf("%s\n", &inputVal)
	if strings.Compare(inputVal, reverseString(inputVal)) != 0 {
		fmt.Println("Entered value is not palindrome")
	} else {
		fmt.Println("Its a palindrome")
	}
}
