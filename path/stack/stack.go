package stack

type Stack struct {
	top  *Node
	size int
}

type Node struct {
	value interface{}
	next  *Node
}

func (s *Stack) Len() int {
	return s.size
}

func (s *Stack) Push(element interface{}) {
	s.top = &Node{element, s.top}
	s.size++
}

func (s *Stack) Pop(element interface{}) {
	if s.size > 0 {
		element, s.top = s.top.value, s.top.next
		s.size--
		return
	}
}

func (s *Stack) Reverse() []int {

}
