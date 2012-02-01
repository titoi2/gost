package main

type StackElement struct {
	element interface{}
}

const StackSize = 100

type Stack struct {
	stackPointer int
	stackMax     int
	stack        []StackElement
}

func (s *Stack) Init(size int) {
	s.stackPointer = 0
	s.stackMax = size
	s.stack = make([]StackElement, size)
}

func (s *Stack) Push(elem interface{}) int {
	if s.stackPointer >= s.stackMax {
		// stack over
		return -1
	}
	s.stack[s.stackPointer].element = elem
	s.stackPointer++
	return 0
}

func (s *Stack) Pop() interface{} {
	if s.stackPointer <= 0 {
		// stack under
		return nil
	}
	s.stackPointer--
	return s.stack[s.stackPointer].element
}
