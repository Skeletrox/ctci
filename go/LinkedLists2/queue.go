package main

// Queue stores the queue for the values
type Queue struct {
	Values   []int
	Capacity int
}

// Push pushes a value to the queue if not full yet. If full, it pops the first element and then pushes
func (q *Queue) Push(value int) bool {
	if len(q.Values) == q.Capacity {
		_, ok := q.Pop()
		if !ok {
			return false
		}
	}
	q.Values = append(q.Values, value)
	return true
}

// Pop removes the first element from the queue, if any element exists
func (q *Queue) Pop() (int, bool) {
	if len(q.Values) == 0 {
		return 0, false
	}
	returnable := q.Values[0]
	q.Values = q.Values[1:]
	return returnable, true
}
