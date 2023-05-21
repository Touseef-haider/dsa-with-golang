package stack

import "fmt"

func push() {

}

type Stack struct {
	stack []int
	top   int
}

func (s *Stack) push(value int) {
	s.stack = append(s.stack, value)
	s.top = s.top + 1
}

func (s *Stack) pop() {
	if len(s.stack) <= 0 {
		fmt.Println("Stack is empty")
	} else {

		s.stack = s.stack[:s.top-1]
		s.top = s.top - 1
	}
}

func (s *Stack) length() int {
	return len(s.stack)
}

func StackDS() {
	var stack Stack = Stack{}
	stack.push(3)
	stack.push(6)
	stack.push(9)
	stack.push(97)
	stack.push(75)
	stack.pop()
	fmt.Println(stack.length())
	fmt.Println(stack)
}
