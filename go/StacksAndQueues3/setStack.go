package main

type SetOfStacks struct {
	StackCapacity int
	Stacks        []*Stack
}

type Stack struct {
	Capacity int
	Elements []int
}

func (s *Stack) Push(value int) bool {
	if s.Elements == nil {
		s.Elements = make([]int, 0)

	} else if len(s.Elements) == s.Capacity {
		return false
	}
	s.Elements = append(s.Elements, value)
	return true
}

func (s *Stack) Pop() (int, bool) {
	if len(s.Elements) == 0 {
		return 0, false
	}
	// all len() methods assume negligible O(1) access times, use a counter if speed is needed
	returnable := s.Elements[len(s.Elements)-1]
	s.Elements = s.Elements[:len(s.Elements)-1]
	return returnable, true
}

func (s *SetOfStacks) Push(value int) {
	if len(s.Stacks) == 0 {
		s.Stacks = make([]*Stack, 0)
		newStack := &Stack{Capacity: s.StackCapacity}
		s.Stacks = append(s.Stacks, newStack)
	}
	candidateStack := s.Stacks[len(s.Stacks)-1]
	if !candidateStack.Push(value) {
		// create new stack
		s2 := Stack{Capacity: s.StackCapacity}
		s2.Push(value)
		s.Stacks = append(s.Stacks, &s2)
	}
}

func (s *SetOfStacks) Pop() (int, bool) {
	if len(s.Stacks) == 0 {
		return 0, false
	}
	candidateStack := s.Stacks[len(s.Stacks)-1]
	val, _ := candidateStack.Pop()
	if len(candidateStack.Elements) == 0 {
		s.Stacks = s.Stacks[:len(s.Stacks)-1]
	}
	return val, true
}

func (s *SetOfStacks) PopAt(index int) (int, bool) {
	if index >= len(s.Stacks) || len(s.Stacks[index].Elements) == 0 {
		return 0, false
	}
	val, ok := s.Stacks[index].Pop()
	// Don't delete empty stacks in the middle
	return val, ok
}
