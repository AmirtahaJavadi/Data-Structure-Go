package main

import "fmt"

func main() {
	stack := Stack{}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	fmt.Println("Stack contents:")
	stack.Print()

	top, err := stack.Peek()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Top of the stack:", top)
	}

	popped, err := stack.Pop()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Popped from stack:", popped)
	}

	fmt.Println("Stack contents after popping:")
	stack.Print()

	fmt.Println("Size of the stack:", stack.Size())
	stack.Push(5)
	fmt.Println("Size of the stack:", stack.Size())
	stack.Print()
}
