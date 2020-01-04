package main

import "fmt"

// Node is a basic Linked List Node
type Node struct {
	Value int
	Next  *Node
}

func helper(head *Node, k int) []int {
	queue := Queue{Capacity: k}
	current := head
	for current != nil {
		queue.Push(current.Value)
		current = current.Next
	}
	return queue.Values
}

func main() {
	linkedList := [15]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	vanguard := &Node{0, nil}
	current := vanguard
	for _, c := range linkedList {
		new := &Node{c, nil}
		current.Next = new
		current = current.Next
	}
	result := helper(vanguard.Next, 8)
	for _, x := range result {
		fmt.Printf("%d ", x)
	}
}
