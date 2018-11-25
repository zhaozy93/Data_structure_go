package stack

import (
	"common"
)

type Stack struct {
	stack *common.Pipe
}

func NewStack(cap int64) *Stack {
	return &Stack{stack: common.NewPipe(cap)}
}

func (stack *Stack) GetLen() int64 {
	return stack.stack.GetLen()
}

func (stack *Stack) Empty() bool {
	return stack.stack.Empty()
}

func (stack *Stack) Push(data interface{}) (*Stack, error) {
	_, err := stack.stack.Push(data)
	return stack, err
}
func (stack *Stack) Pop() (interface{}, error) {
	return stack.stack.Pop()
}
