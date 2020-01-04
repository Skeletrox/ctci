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

func helper(head *Node, value int) *Node {
	smallerVanguard := &Node{0, nil}
	largerVanguard := &Node{0, nil}
	smallPtr := smallerVanguard
	largePtr := largerVanguard
	currPtr := head
	var pivotPtr *Node = nil
	for currPtr != nil {
		if currPtr.Value < value {
			smallPtr.Next = currPtr
			smallPtr = smallPtr.Next
		} else if pivotPtr == nil && currPtr.Value == value {
			pivotPtr = currPtr
		} else {
			largePtr.Next = currPtr
			largePtr = largePtr.Next
		}
		currPtr = currPtr.Next
	}
	smallPtr.Next = pivotPtr
	pivotPtr.Next = largerVanguard.Next
	largePtr.Next = nil

	return smallerVanguard.Next
}

func main() {
	linkedList := [20]int{20, 19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	vanguard := &Node{0, nil}
	current := vanguard
	for _, c := range linkedList {
		new := &Node{c, nil}
		current.Next = new
		current = current.Next
	}
	rand.Seed(time.Now().UnixNano())
	pivot := rand.Int()%20 + 1
	head := helper(vanguard.Next, pivot)
	fmt.Printf("Position: %d\n", pivot)
	curr := head
	for curr != nil {
		fmt.Printf("%d ", curr.Value)
		curr = curr.Next
	}
}
