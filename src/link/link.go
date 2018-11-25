package link

import (
	"common"
)

func NewLink(cap int64) *common.Pipe {
	return common.NewPipe(cap)
}
