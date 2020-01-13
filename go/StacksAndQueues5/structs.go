package main

type SortedStack struct {
	Main, Temp *Stack
}
type Stack struct {
	Values []int
}

func (s *Stack) Push(value int) {
	if s.Values == nil {
		s.Values = make([]int, 0)
	}
	s.Values = append(s.Values, value)
}

func (s *Stack) Pop() (int, bool) {
	if s.Values == nil || len(s.Values) == 0 {
		return 0, false
	}
	returnable := s.Values[len(s.Values)-1]
	s.Values = s.Values[:len(s.Values)-1]
	return returnable, true
}

func (s *Stack) Peek() (int, bool) {
	if s.Values == nil || len(s.Values) == 0 {
		return 0, false
	}
	return s.Values[len(s.Values)-1], true
}

func (s *SortedStack) Peek() (int, bool) {
	if s.Main == nil {
		return 0, false
	}
	return s.Main.Peek()
}

func (s *SortedStack) Push(value int) {
	if s.Main == nil {
		s.Main = &Stack{}
		s.Temp = &Stack{}
	}
	for len(s.Main.Values) > 0 {
		val, _ := s.Main.Peek()
		if val < value {
			val, _ = s.Main.Pop()
			s.Temp.Push(val)
			continue
		}
		break
	}
	s.Main.Push(value)
	for len(s.Temp.Values) > 0 {
		res, _ := s.Temp.Pop()
		s.Main.Push(res)
	}
}

func (s *SortedStack) Pop() (int, bool) {
	return s.Main.Pop()
}
