package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) push(a int) []int {
	s.items = append(s.items, a)
	return s.items
}

func (s *Stack) pop() int {
	popedItem := s.items[len(s.items)-1]
	s.items = s.items[0:(len(s.items) - 1)]
	return popedItem
}

func (s *Stack) size() int {
	return len(s.items)
}

func main() {
	stack := &Stack{}
	stack.items = []int{1, 2, 3, 4}

	stack.push(5)
	fmt.Println(stack.items)
	stack.pop()
	fmt.Println(stack.items)
	fmt.Println(stack.size())
}
