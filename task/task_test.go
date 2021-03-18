package task

import (
	"context"
	"errors"
	"fmt"
	"runtime"
	"testing"
	"time"
)

type timeout struct {
}

func (t timeout) CallBack(p []byte) {
	//println("run callback")
	println(string(p), "执行成功")
}

func (t timeout) Payload() ([]byte, error) {
	println("run payload timeout.")
	time.Sleep(4 * time.Second)
	return []byte(`超时任务`), nil
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
		NewLink(timeout{}),
		NewLink(err{name: "one"}),
		NewLink(normal{}),
		NewLink(err{name: "two"}),
		NewLink(err{name: "three"}),
		NewLink(normal{}),
	)
	task.Each()
	//time.Sleep(4 * 6 * time.Second)
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
