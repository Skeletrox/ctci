package main

import "fmt"

func main() {
	queue := Queue{}
	for i := 0; i < 20; i++ {
		queue.Enqueue(i)
	}
	for i := 0; i < 20; i++ {
		val, _ := queue.Dequeue()
		fmt.Println(val)
	}
}
