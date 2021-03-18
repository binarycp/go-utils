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

type timeout struct {
}

func (t timeout) CallBack(p []byte) {
	//println("run callback")
	println(string(p), "执行成功")
}

func (t timeout) Payload() ([]byte, error) {
	println("run payload timeout.")
	time.Sleep(400 * time.Millisecond)
	return []byte(`超时任务`), errors.New("run timeout")
}

type normal struct {
	timeout
}

func (n normal) Payload() ([]byte, error) {
	println("run payload")
	return []byte(`正常任务`), nil
}

type err struct {
	timeout
	name string
}

func (e err) Payload() ([]byte, error) {
	println(e.name, "run payload loop.")
	return []byte(e.name + `错误任务`), errors.New("run continue")
}

func TestTask_Each(t1 *testing.T) {
	t1.Helper()
	task := NewTask(3 * time.Second)
	task.Add(
		NewLink(3*time.Millisecond, timeout{}),
		NewLink(3*time.Millisecond, err{name: "one"}),
		NewLink(3*time.Millisecond, normal{}),
		NewLink(3*time.Millisecond, err{name: "two"}),
		NewLink(3*time.Millisecond, err{name: "three"}),
		NewLink(3*time.Millisecond, normal{}),
	)
	go task.Each()
	time.Sleep(4 * 6 * time.Second)
	t1.Log(runtime.NumGoroutine())
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

func TestChan(t *testing.T) {
	c := make(chan struct{}, 2)
	println(cap(c), len(c))
	c <- struct{}{}
	println(cap(c), len(c))
}
