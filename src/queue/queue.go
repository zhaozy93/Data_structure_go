package queue

import (
	"common"
)

type Queue struct {
	queue *common.Pipe
}

func NewQueue(cap int64) *Queue {
	return &Queue{queue: common.NewPipe(cap)}
}

func (queue *Queue) GetLen() int64 {
	return queue.queue.GetLen()
}

func (queue *Queue) Empty() bool {
	return queue.queue.Empty()
}

func (queue *Queue) Push(data interface{}) (*Queue, error) {
	_, err := queue.queue.Push(data)
	return queue, err
}

func (queue *Queue) Shift() (interface{}, error) {
	return queue.queue.Shift()
}
