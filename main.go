package main

import (
	"data_structure/filo"
	"data_structure/stack"
	"fmt"
	"log"
)

func main() {
	log.Printf("SELECT AN OPTION BELOW TO SEE THE EXAMPLE RUNING\n")
	log.Printf("\n1 - STACK (FILO)\n2 - QUEUE (FIFO)")

	var option string
	_, err := fmt.Scanln(&option)
	if err != nil {
		log.Println(err.Error())
		return
	}

	switch option {
	case "1":
		RunStack()

	case "2":
		RunQueue()

	default:
		log.Println("The option selected is invalid")
	}
}

func RunStack() {
	stack := stack.NewStack()

	stack.Push("first")
	stack.Push("second")
	stack.Push("third")

	var value string
	var err error

	for {
		log.Println("Complete stack")
		stack.Print()

		value, err = stack.Pop()

		if err != nil {
			log.Println(err.Error())
			break
		}

		log.Printf("\nRemoved value: %s\n", value)
	}
}

func RunQueue() {
	queue := filo.NewQueue()

	queue.Push("first")
	queue.Push("second")
	queue.Push("third")

	for {
		log.Println("Complete queue")
		queue.Print()

		value, err := queue.Pop()

		if err != nil {
			log.Println(err.Error())
			break
		}

		log.Printf("\nRemoved value: %s\n", value)
	}
}
