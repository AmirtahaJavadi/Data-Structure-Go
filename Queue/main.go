package main

import "fmt"

func main() {
	q := Queue{}
	q.Enqueue(5)
	q.Enqueue(10)
	q.Enqueue(15)
	q.print()
	q.Dequeue()
	q.print()
	fmt.Println(q.IsEmpty())
}
