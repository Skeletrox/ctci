package main

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
