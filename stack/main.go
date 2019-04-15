package main

import "errors"
import "fmt"

type Stack struct {
	slice []int
}

func (stack *Stack) Push(n int) {
	stack.slice = append(stack.slice, n)
}

func (stack *Stack) Peek() (int, error) {
	if len(stack.slice) == 0 {
		return 0, errors.New("Không tìm thấy phần tử")
	}

	lastItem := stack.slice[len(stack.slice)-1]
	return lastItem, nil
}

func (stack *Stack) Pop() (int, error) {
	lastItem, err := stack.Peek()
	if err != nil {
		return 0, err
	}

	stack.slice = stack.slice[0:len(stack.slice)-1]
	return lastItem, nil
}

func main() {
	stack := Stack{}
	var lastItem int

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)

	fmt.Println(stack)

	lastItem, _ = stack.Pop()
	fmt.Println(lastItem)

	lastItem, _ = stack.Pop()
	fmt.Println(lastItem)
	

	fmt.Println(stack)
}
