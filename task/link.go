package task

type Procedure interface {
	CallBack([]byte)
	Payload() ([]byte, error)
}

// 定时执行任务
// 具备超时时间
type Link struct {
	Next *Link
	Procedure
}

// 实例化链表
func NewLink(p Procedure) *Link {
	return &Link{
		Next:      nil,
		Procedure: p,
	}
}
