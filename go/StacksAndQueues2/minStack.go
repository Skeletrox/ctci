package main

type MinStackStruct struct {
	Stack, MinStack []int
}

func (m *MinStackStruct) Push(val int) {
	if m.Stack == nil {
		m.Stack = make([]int, 0)
		m.MinStack = make([]int, 0)
		m.Stack = append(m.Stack, val)
		m.MinStack = append(m.Stack, val)
	} else {
		m.Stack = append(m.Stack, val)
		if val <= m.MinStack[len(m.MinStack)-1] {
			m.MinStack = append(m.MinStack, val)
		}
	}

}

func (m *MinStackStruct) Pop() (int, bool) {
	if len(m.Stack) == 0 {
		return 0, false
	}
	returnable := m.Stack[len(m.Stack)-1]
	m.Stack = m.Stack[:len(m.Stack)-1]
	if returnable <= m.MinStack[len(m.MinStack)-1] {
		m.MinStack = m.MinStack[:len(m.MinStack)-1]
	}
	return returnable, true
}

func (m *MinStackStruct) Min() int {
	if len(m.MinStack) == 0 {
		return 0
	}
	return m.MinStack[len(m.MinStack)-1]
}
