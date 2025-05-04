package fizzBuzz

type Stack struct {
	elements []int
	size     int
}

func (s *Stack) push(i int) {
	s.elements = append([]int{i}, s.elements...)
}

func (s *Stack) Size() int {
	return len(s.elements)
}

func (s *Stack) pop() int {
	item := s.elements[:len(s.elements)]
	s.elements = append(s.elements[1:], s.elements[:0]...)
	return item[0]
}
