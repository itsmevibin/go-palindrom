package main

import "fmt"

func reserseString(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func main() {
	fmt.Println(reserseString("hello world"))
}
