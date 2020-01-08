package main

import "fmt"

func main() {
	ms := MinStackStruct{}
	ms.Push(5)
	ms.Push(3)
	min := ms.Min()
	fmt.Println("Min is", min)
	ms.Push(4)
	min = ms.Min()
	fmt.Println("Min is", min)
	ms.Push(2)
	min = ms.Min()
	fmt.Println("Min is", min)
	popped, _ := ms.Pop()
	min = ms.Min()
	fmt.Println("Popped is", popped, "Min is", min)
}
