package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func helper(head *Node) bool {

	slowRunner := head
	fastRunner := head

	stack := make([]int, 0)

	for fastRunner != nil && fastRunner.Next != nil {
		stack = append(stack, slowRunner.Value)
		slowRunner = slowRunner.Next
		fastRunner = fastRunner.Next.Next
	}

	if fastRunner == nil {
		// Even number of nodes, slowRunner is at the "middle-left" node
		for len(stack) != 0 {
			val := stack[len(stack)-1]
			if slowRunner == nil || slowRunner.Value != val {
				return false
			}
			slowRunner = slowRunner.Next
			stack = stack[:len(stack)-1]
		}
	} else {
		// Odd number of nodes, slowRunner is currently at the middle node
		slowRunner = slowRunner.Next
		for len(stack) != 0 {
			val := stack[len(stack)-1]
			if slowRunner == nil || slowRunner.Value != val {
				return false
			}
			slowRunner = slowRunner.Next
			stack = stack[:len(stack)-1]
		}
	}

	return true
}

func main() {
	numList := [7]int{1, 2, 3, 4, 3, 2, 1}
	vanguard := &Node{0, nil}
	ptr := vanguard
	for _, c := range numList {
		curr := &Node{c, nil}
		ptr.Next = curr
		ptr = ptr.Next
	}
	result := helper(vanguard.Next)
	fmt.Println("Result is", result)
}
