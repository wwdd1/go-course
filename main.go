package main

import (
	"errors"
	"fmt"
)

type Stack []int

func (stack *Stack) IsEmpty() bool {
	return len(*stack) == 0
}

func (stack *Stack) Push(value int) {
	*stack = append(*stack, value)
}

func (stack *Stack) Pop() (error, int) {
	if stack.IsEmpty() {
		return errors.New("Stack is empty"), 0
	}

	lastIndex := len(*stack) - 1
	value := (*stack)[lastIndex]
	*stack = (*stack)[:lastIndex]

	return nil, value
}

func main() {
	var stack Stack
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	err, value := stack.Pop()
	if err == nil {
		fmt.Printf("We got: %v \n", value)
	}
	fmt.Printf("Stack is empty: %v \n", stack.IsEmpty())
	fmt.Printf("Final stack is %v \n", stack)

}
