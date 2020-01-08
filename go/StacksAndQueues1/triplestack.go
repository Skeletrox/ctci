package main

type TripleStack struct {
	Container  []int
	StackHeads [3]int
	StackSizes [3]int
}

func (t *TripleStack) Init() {
	t.Container = make([]int, 3)
	t.StackHeads = [3]int{0, 1, 2}
	t.StackSizes = [3]int{0, 0, 0}
}

func (t *TripleStack) Push(val, stackNum int) {
	loc := t.StackHeads[stackNum]
	t.Container = append(t.Container, 0)
	copy(t.Container[loc+1:], t.Container[loc:])
	t.Container[loc] = val
	t.StackSizes[stackNum]++
	for i := stackNum + 1; i < 3; i++ {
		t.StackHeads[i]++
	}
}

func (t *TripleStack) Pop(stackNum int) (int, bool) {

	if t.StackSizes[stackNum] == 0 {
		return 0, false
	}
	loc := t.StackHeads[stackNum]
	returnable := t.Container[loc]
	t.StackSizes[stackNum]--
	copy(t.Container[loc:], t.Container[loc+1:])
	for i := stackNum + 1; i < 3; i++ {
		t.StackHeads[i]--
	}
	return returnable, true
}
