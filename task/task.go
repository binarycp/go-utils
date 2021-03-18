package task

import "time"

type Task struct {
	interval time.Duration
	link     *Link
	stop     bool
}

// 实例化任务
func NewTask(interval time.Duration) *Task {
	return &Task{
		interval: interval,
		link:     nil,
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
	close(t)
	each(t)
	for t.link != nil {
		time.Sleep(t.interval)
		each(t)
	}
}

// 终止任务
func (t *Task) Stop() {
	t.stop = true
}

// 遍历任务链表
func each(task *Task) {
	t := task.link
	if t == nil {
		return
	}
	defer close(task)

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
			if err == nil {
				t.CallBack(payload)
			}
		}()
		select {
		case <-done:
			if err != nil {
				list = append(list, t)
			}
		case <-time.After(task.interval):
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

func close(task *Task) {
	if task.stop {
		task.link = nil
	}
}
