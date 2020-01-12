package main

import "fmt"

func main() {
	capacity := 5
	stackSet := SetOfStacks{StackCapacity: capacity}
	for i := 0; i < 20; i++ {
		stackSet.Push(i + 1)
	}
	fmt.Println("Number of stacks:", len(stackSet.Stacks))
	for i := 0; i < 20; i++ {
		val, ok := stackSet.Pop()
		if !ok {
			fmt.Println("Could not pop!")
		} else {
			fmt.Println(val)
		}
	}
	for i := 0; i < 20; i++ {
		stackSet.Push(i + 1)
	}
	for i := 0; i < 6; i++ {
		val, ok := stackSet.PopAt(1)
		if !ok {
			fmt.Println("Could not pop!")
		} else {
			fmt.Println(val)
		}
	}
}
