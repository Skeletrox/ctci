package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Node struct {
	Value int
	Next  *Node
}

func helper(node *Node) {
	node.Value = node.Next.Value
	node.Next = node.Next.Next
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
	now := vanguard.Next
	rand.Seed(time.Now().UnixNano())
	position := rand.Int()%13 + 1
	for i := 0; i < position; i++ {
		now = now.Next
	}
	helper(now)
	fmt.Printf("Position: %d\n", position)
	curr := vanguard.Next
	for curr != nil {
		fmt.Printf("%d ", curr.Value)
		curr = curr.Next
	}
}
