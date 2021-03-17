package task

import "time"

type Task struct {
	interval time.Duration
	link     *Link
}

// 定时执行任务
// 具备超时时间
type Link struct {
	Timeout  time.Duration
	Next     *Link
	Callback func([]byte)
	Payload  func() ([]byte, error)
}

// 实例化任务
func NewTask(interval time.Duration) *Task {
	return &Task{
		interval: interval,
		link:     nil,
	}
}

// 实例化链表
func NewLink(timeout time.Duration, callback func([]byte), payload func() ([]byte, error)) *Link {
	return &Link{
		Timeout:  timeout,
		Next:     nil,
		Callback: callback,
		Payload:  payload,
	}
}

// 向链表添加元素
func (t *Task) Add(links ...*Link) {
	for index := len(links) - 1; index >= 0; index-- {
		links[index].Next, t.link = t.link, links[index]
	}
}

// 遍历任务链表
func (t *Task) Each() {
	each(t)
	for t.link != nil {
		time.Sleep(t.interval)
		each(t)
	}
}

// 遍历任务链表
func each(task *Task) {
	t := task.link
	if t == nil {
		return
	}
	var payload []byte
	var err error
	list := make([]*Link, 0)
	ticker := time.NewTicker(task.interval)
	defer ticker.Stop()
	for ; true; <-ticker.C {
		done := make(chan struct{}, 1)
		go func() {
			payload, err = t.Payload()
			done <- struct{}{}
		}()
		select {
		case <-done:
			if err == nil {
				if t.Callback != nil {
					t.Callback(payload)
				}
			} else {
				list = append(list, t)
			}
		case <-time.After(t.Timeout):
			list = append(list, t)
		}

		t = t.Next
		if t == nil {
			break
		}
	}

	var ret *Link
	for index := len(list) - 1; index >= 0; index-- {
		list[index].Next, ret = ret, list[index]
	}
	task.link = ret
}
