package stack_queue

type Stack struct {
	stack []any
}

func NewStack() *Stack {
	return &Stack{
		stack: []any{},
	}
}

func (s *Stack) Push(element any) {
	s.stack = append(s.stack, element)
}

func (s *Stack) Pop() any {
	data := s.stack[len(s.stack)-1]
	s.stack = append(s.stack[:len(s.stack)-1], s.stack[len(s.stack):]...)
	return data
}

type Queue struct {
	queue []any
}

func NewQueue() *Queue {
	return &Queue{
		queue: []any{},
	}
}

func (q *Queue) Push(element any) {
	q.queue = append(q.queue, element)
}

func (q *Queue) PopLeft() any {
	data := q.queue[0]
	q.queue = append(q.queue[:0], q.queue[1:]...)
	return data
}
