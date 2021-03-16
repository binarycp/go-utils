package task

import (
	"context"
	"errors"
	"fmt"
	"runtime"
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

	errPayload1 = func() ([]byte, error) {
		println("run payload loop1.")
		return []byte(`错误任务1`), errors.New("run continue")
	}

	errPayload2 = func() ([]byte, error) {
		println("run payload loop2.")
		return []byte(`错误任务2`), errors.New("run continue")
	}

	timeoutPayload = func() ([]byte, error) {
		println("run payload timeout.")
		time.Sleep(400 * time.Millisecond)
		return []byte(`超时任务`), errors.New("run timeout")
	}
)

func TestTask_Each(t1 *testing.T) {
	task := NewTask(3 * time.Second)
	task.Add(NewLink(3*time.Millisecond, callBack, timeoutPayload),
		NewLink(3*time.Millisecond, callBack, errPayload),
		NewLink(3*time.Millisecond, callBack, normalPayload),
		NewLink(3*time.Millisecond, callBack, errPayload1),
		NewLink(3*time.Millisecond, callBack, errPayload2),
		NewLink(3*time.Millisecond, callBack, normalPayload),
	)
	task.Each()
}

func TestGoroutine(t *testing.T) {
	t.Helper()
	timeout, _ := context.WithTimeout(context.Background(), time.Second)
	for i := 0; i < 1; i++ {
		go loop(timeout)
	}
	time.Sleep(time.Second * 2)
	t.Log(runtime.NumGoroutine())
}

func loop(c context.Context) {
	t := time.Now()
	for {
		select {
		case <-c.Done():
			fmt.Println(time.Since(t))
			return
		default:
		}
	}
}
