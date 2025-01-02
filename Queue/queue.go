package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	items []interface{}
}

func (q *Queue) Enqueue(item interface{}) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() (interface{}, error) {
	if len(q.items) == 0 {
		return nil, errors.New("Queue is Empty")
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, nil
}

func (q *Queue) Size() int {
	return len(q.items)
}

func (q *Queue) Peek() (interface{}, error) {
	if len(q.items) == 0 {
		return nil, errors.New("queue is empty")
	}
	return q.items[0], nil
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue) print() {
	for _, v := range (*q).items {
		fmt.Println(v)

	}
	fmt.Println("")
}
