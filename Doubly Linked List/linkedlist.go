package main

import "fmt"

type Linkedlist struct {
	first *Node
}

func (l *Linkedlist) add(n int) {
	var temp *Node
	newNode := &Node{data: n}
	if l.first == nil {
		l.first = newNode // If the list is empty, set the first node
		return
	}

	current := l.first
	fmt.Println("current is: ", current)
	temp = current
	for current.next != nil {
		current = current.next
		temp = current
		fmt.Println("current is(in for): ", current)
		if current == nil {
			break
		}
	}
	// Link the new node to the end of the list
	current.next = newNode
	newNode.prev = temp
	fmt.Println("new Node : ", current)
}

func (l *Linkedlist) reverse() {
	var current *Node
	q := l.first
	for l.first != nil {
		l.first = l.first.next
		q.next = current
		current = q
		q = l.first
	}
	l.first = current

}

func (l *Linkedlist) print() {
	fmt.Println("printing...")
	next := l.first
	for next != nil {
		fmt.Println(next)
		fmt.Print(" -> ")
		next = next.next
	}
	fmt.Printf("\n")
}
