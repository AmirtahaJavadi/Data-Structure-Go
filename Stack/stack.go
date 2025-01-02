package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	items []interface{}
}

func (s *Stack) Peek() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	return s.items[len(s.items)-1], nil
}

func (s *Stack) Push(value interface{}) {
	s.items = append(s.items, value)
}
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}
func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]
	return item, nil
}

func (s *Stack) Size() int {
	return len(s.items)
}

func (s *Stack) Print() {
	for i := len(s.items) - 1; i >= 0; i-- {
		fmt.Println(s.items[i])
	}
}
