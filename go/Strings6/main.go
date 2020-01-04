package main

import (
	"fmt"
	"strings"
)

func helper(input string) string {
	var returnable strings.Builder
	char := rune(input[0])
	count := 1
	for _, c := range input[1:] {
		if c != char {
			returnable.WriteRune(char)
			returnable.WriteString(fmt.Sprintf("%d", count))
			char = c
			count = 1
		} else {
			count++
		}
	}
	returnable.WriteRune(char)
	returnable.WriteString(fmt.Sprintf("%d", count))
	res := returnable.String()
	if len(res) < len(input) {
		return res
	}
	return input
}

func main() {
	inputs := [4]string{"abcdefgh", "aaaaaaaaaa", "aabbbcdddeeef", "xxxxxx"}
	for _, i := range inputs {
		res := helper(i)
		fmt.Println("Original:", i, "Compressed:", res)
	}
}
