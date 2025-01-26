package filo

import (
	"errors"
	"log"
)

type node struct {
	value string
	next  *node
	prev  *node
}

type Filo struct {
	head *node
	last *node
}

func NewQueue() *Filo {
	return &Filo{}
}

// Push add a value to the end of the queue
func (f *Filo) Push(value string) {
	newNode := &node{
		value: value,
	}

	if f.head == nil {
		f.head = newNode
		f.last = newNode
		return
	}

	newNode.prev = f.last
	f.last.next = newNode
	f.last = newNode
}

// Pop returns the first value of the queue removing the value
func (f *Filo) Pop() (string, error) {
	if f.head == nil {
		return "", errors.New("THE QUEUE IS EMPTY")
	}

	value := f.head.value

	if f.head == f.last {
		f.last = f.head.next
	}

	f.head = f.head.next
	return value, nil
}

// Print show all itens in the queue without removing any
func (f *Filo) Print() {
	currentNode := f.head
	for currentNode != nil {
		log.Println(currentNode.value)
		currentNode = currentNode.next
	}
}

// GetLast returns the last inserted value of the without removing
func (f *Filo) GetLast() (string, error) {
	if f.last == nil {
		return "", errors.New("THE QUEUE IS EMPTY")
	}

	return f.last.value, nil
}

// GetHead retuns the first value in the queue without removing
func (f *Filo) GetHead() (string, error) {
	if f.head == nil {
		return "", errors.New("THE QUEUE IS EMPTY")
	}

	return f.head.value, nil
}
