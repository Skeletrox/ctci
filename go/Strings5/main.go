package main

import (
	"fmt"
	"math"
)

func helper(one, two string) bool {
	if math.Abs(float64(len(one)-len(two))) > 1 {
		return false
	}
	var small, big string
	if len(one) < len(two) {
		small = one
		big = two
	} else {
		small = two
		big = one
	}
	changeFound := false
	for i, c := range small {
		if c != rune(big[i]) {
			if !changeFound {
				changeFound = true
			} else {
				return false
			}
		}
	}
	return true
}

func main() {
	inputs1 := [4]string{"hello", "borscht", "ngazargamu", "fordgt"}
	inputs2 := [4]string{"hell", "porsche", "xyz", "bordgt"}
	for i := 0; i < 4; i++ {
		res := helper(inputs1[i], inputs2[i])
		fmt.Println("String 1:", inputs1[i], "String 2:", inputs2[i], "Result:", res)
	}
}
