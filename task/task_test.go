package task

import (
	"errors"
	"testing"
	"time"
)

var (
	callBack = func(p []byte) {
		//println("run callback")
		println(string(p), "执行成功")
	}

	normalPayload = func() ([]byte, error) {
		println("run payload")
		return []byte(`正常任务`), nil
	}

	errPayload = func() ([]byte, error) {
		println("run payload loop.")
		return []byte(`错误任务`), errors.New("run continue")
	}
)

func TestTask_Each(t1 *testing.T) {
	task := NewTask(3 * time.Second)
	task.Add(NewLink(3*time.Millisecond, callBack, normalPayload), NewLink(3*time.Millisecond, callBack, errPayload), NewLink(3*time.Millisecond, callBack, normalPayload))
	task.Each()
}
