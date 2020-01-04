package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func helper(head *Node) *Node {
	seen := make(map[int]bool)
	current := head

	for current != nil && current.Next != nil {
		seen[current.Value] = true
		if seen[current.Next.Value] {
			current.Next = current.Next.Next
		}
		current = current.Next
	}
	return head
}

func main() {
	nums := [10]int{1, 2, 3, 3, 4, 5, 6, 7, 8, 9}
	vanguard := &Node{0, nil}
	current := vanguard
	for _, n := range nums {
		x := &Node{n, nil}
		current.Next = x
		current = x
	}
	head := helper(vanguard.Next)
	for head != nil {
		fmt.Printf("%d ", head.Value)
		head = head.Next
	}
}
