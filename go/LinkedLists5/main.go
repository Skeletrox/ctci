package main

import "fmt"

type Node struct {
	Next  *Node
	Value int
}

func helper(oneHead, twoHead *Node) *Node {

	returnableVanguard := &Node{nil, 0}
	returnablePtr := returnableVanguard

	onePtr := oneHead
	twoPtr := twoHead

	carry := 0

	for onePtr != nil && twoPtr != nil {
		val := onePtr.Value + twoPtr.Value + carry
		carry = val / 10
		val %= 10
		current := &Node{nil, val}
		returnablePtr.Next = current
		returnablePtr = returnablePtr.Next
		onePtr = onePtr.Next
		twoPtr = twoPtr.Next
	}

	if onePtr == nil {
		for twoPtr != nil {
			val := twoPtr.Value + carry
			carry = val / 10
			val %= 10
			current := &Node{nil, val}
			returnablePtr.Next = current
			returnablePtr = returnablePtr.Next
			twoPtr = twoPtr.Next
		}
	}

	if twoPtr == nil {
		for onePtr != nil {
			val := onePtr.Value + carry
			carry = val % 10
			val /= 10
			current := &Node{nil, val}
			returnablePtr.Next = current
			returnablePtr = returnablePtr.Next
			onePtr = onePtr.Next
		}
	}

	if carry == 1 {
		curr := &Node{nil, 1}
		returnablePtr.Next = curr
	}

	return returnableVanguard.Next
}

func main() {

	num1 := [3]int{9, 6, 3}
	num2 := [3]int{2, 6, 5}

	num1Vanguard := &Node{nil, 0}
	num2Vanguard := &Node{nil, 0}

	num1Ptr := num1Vanguard
	num2Ptr := num2Vanguard

	for i := range num1 {

		curr1 := &Node{nil, num1[i]}
		curr2 := &Node{nil, num2[i]}

		num1Ptr.Next = curr1
		num2Ptr.Next = curr2

		num1Ptr = num1Ptr.Next
		num2Ptr = num2Ptr.Next
	}

	result := helper(num1Vanguard.Next, num2Vanguard.Next)
	for result != nil {
		fmt.Printf("%d ", result.Value)
		result = result.Next
	}
}
