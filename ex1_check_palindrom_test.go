package main

import "testing"

func TestReverse(t *testing.T) {
	if "hello" != reserseString("olleh") {
		t.Errorf("Reverse string is failed")
	}
}
