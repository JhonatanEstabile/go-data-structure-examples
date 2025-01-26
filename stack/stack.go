package stack

import (
	"errors"
	"log"
)

type stack_node struct {
	value string
	next  *stack_node
}

type Stack struct {
	top *stack_node
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(value string) {
	newNode := &stack_node{
		value: value,
		next:  s.top,
	}

	s.top = newNode
}

func (s *Stack) Pop() (string, error) {
	if s.top == nil {
		return "", errors.New("STACK IS EMPTY")
	}

	value := s.top.value
	s.top = s.top.next

	return value, nil
}

func (s *Stack) Peek() (string, error) {
	if s.top == nil {
		return "", errors.New("STACK IS EMPTY")
	}

	return s.top.value, nil
}

func (s *Stack) Print() {
	mainNode := s.top

	for mainNode != nil {
		log.Println(mainNode.value)
		mainNode = mainNode.next
	}
}
