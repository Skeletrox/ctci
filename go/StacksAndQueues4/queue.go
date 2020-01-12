package main

type Queue struct {
	Stack1, Stack2 *Stack
}

func (q *Queue) Enqueue(value int) {
	if q.Stack1 == nil {
		q.Stack1 = &Stack{}
		q.Stack2 = &Stack{}
	}
	q.Stack1.Push(value)
}

func (q *Queue) Dequeue() (int, bool) {
	if q.Stack1 == nil {
		return 0, false
	}
	for {
		if val, ok := q.Stack1.Pop(); ok {
			q.Stack2.Push(val)
		} else {
			break
		}
	}
	returnable, _ := q.Stack2.Pop()
	for {
		if val, ok := q.Stack2.Pop(); ok {
			q.Stack1.Push(val)
		} else {
			break
		}
	}
	return returnable, true
}
