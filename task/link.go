package task

import "time"

type Procedure interface {
	CallBack([]byte)
	Payload() ([]byte, error)
}

// 定时执行任务
// 具备超时时间
type Link struct {
	Timeout time.Duration
	Next    *Link
	Procedure
}

// 实例化链表
func NewLink(timeout time.Duration, p Procedure) *Link {
	return &Link{
		Timeout:   timeout,
		Next:      nil,
		Procedure: p,
	}
}
