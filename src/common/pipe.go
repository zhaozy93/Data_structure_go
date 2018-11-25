package common

import (
	"errors"
	"math"
	"sync"
)

var ERRTOOLARGE = errors.New("exceed the capacity")
var ERRTOOSHORT = errors.New("no item existed")

type Pipe struct {
	data   *Item
	start  *Item
	end    *Item
	len    int64
	cap    int64
	locker *sync.Mutex
}

func NewPipe(cap int64) *Pipe {
	if cap <= 0 {
		cap = math.MaxInt64
	}
	return &Pipe{
		data:   nil,
		start:  nil,
		end:    nil,
		len:    0,
		cap:    cap,
		locker: new(sync.Mutex),
	}
}

func (pipe *Pipe) Empty() bool {
	pipe.locker.Lock()
	defer pipe.locker.Unlock()
	return pipe.start == nil
}

func (pipe *Pipe) GetLen() int64 {
	pipe.locker.Lock()
	defer pipe.locker.Unlock()
	return pipe.len
}

func (pipe *Pipe) Push(data interface{}) (*Pipe, error) {
	pipe.locker.Lock()
	defer pipe.locker.Unlock()
	if pipe.start == nil {
		item := NewItem(data)
		pipe.data = item
		pipe.start = item
		pipe.end = item
		pipe.len++
		return pipe, nil
	} else if pipe.len > pipe.cap-1 {
		return pipe, ERRTOOLARGE
	}
	item := NewItem(data)
	pipe.end.SetNext(item)
	item.SetPre(pipe.end)
	pipe.end = item
	pipe.len++
	return pipe, nil
}

func (pipe *Pipe) Pop() (interface{}, error) {
	pipe.locker.Lock()
	defer pipe.locker.Unlock()
	if pipe.start == nil {
		return nil, ERRTOOSHORT
	}
	end := pipe.end
	pipe.len--
	if pipe.len == 0 {
		pipe.data = nil
		pipe.start = nil
		pipe.end = nil
	} else {
		last := end.GetPre()
		last.SetNext(nil)
		pipe.end = last
	}
	return end.GetValue(), nil
}

func (pipe *Pipe) Unshift(data interface{}) (*Pipe, error) {
	pipe.locker.Lock()
	defer pipe.locker.Unlock()
	if pipe.start == nil {
		item := NewItem(data)
		pipe.data = item
		pipe.start = item
		pipe.end = item
		pipe.len++
		return pipe, nil
	} else if pipe.len > pipe.cap-1 {
		return pipe, ERRTOOLARGE
	}
	item := NewItem(data)
	pipe.start.SetPre(item)
	item.SetNext(pipe.start)
	pipe.start = item
	pipe.len++
	return pipe, nil
}

func (pipe *Pipe) Shift() (interface{}, error) {
	pipe.locker.Lock()
	defer pipe.locker.Unlock()
	if pipe.start == nil {
		return nil, ERRTOOSHORT
	}
	start := pipe.start
	pipe.len--
	if pipe.len == 0 {
		pipe.data = nil
		pipe.start = nil
		pipe.end = nil
	} else {
		first := start.GetNext()
		first.SetPre(nil)
		pipe.start = first
	}
	return start.GetValue(), nil
}
