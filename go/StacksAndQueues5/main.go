package main

import "fmt"

func main() {
	sStack := SortedStack{}
	values := [10]int{10, 6, 4, 9, 8, 1, 7, 5, 2, 3}
	for _, v := range values {
		sStack.Push(v)
	}
	for {
		val, ok := sStack.Pop()
		if !ok {
			break
		}
		fmt.Println(val)
	}
}
