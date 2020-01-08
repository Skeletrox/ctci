package main

import "fmt"

func main() {
	t := TripleStack{}
	t.Init()
	t.Push(2, 0)
	t.Push(1, 1)
	t.Push(3, 2)
	t.Push(4, 0)
	var ok bool
	var v int
	v, ok = t.Pop(0)
	if !ok {
		fmt.Println("Shouldn't be here")
		return
	} else {
		fmt.Println(v)
	}
	v, ok = t.Pop(0)
	if !ok {
		fmt.Println("Shouldn't be here")
		return
	} else {
		fmt.Println(v)
	}
	v, ok = t.Pop(0)
	if ok {
		fmt.Println("Shouldn't be here")
		return
	} else {
		fmt.Println("Cannot pop!")
	}
}
