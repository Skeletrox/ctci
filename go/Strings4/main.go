package main

import (
	"fmt"
	"strings"
)

func helper(input string) bool {
	// Assuming case insensitivity
	var x uint32
	fixed := strings.ToLower(input)
	for _, c := range fixed {
		if c >= 'a' && c <= 'z' {
			x ^= 1 << (c - 'a')
		}
	}
	return x == 0 || x&(x-1) == 0
}

func main() {
	words := [4]string{"Cat Taco", "Surrender", "Ford GT", "kayak"}
	for _, w := range words {
		res := helper(w)
		fmt.Println("Word:", w, "is a palindrome permuation?", res)
	}
}
