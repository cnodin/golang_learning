package select_test

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Second * 1)

	return "Done"
}

func AsyncService() chan string {
	retCh := make(chan string, 1)
	go func() {
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret
		fmt.Println("service exited")
	}()

	return retCh
}

func TestSelect(t *testing.T) {
	select {
		case ret := <- AsyncService():
			t.Log(ret)
		case <-time.After(time.Second * 10):
			t.Error("timeout")
	}
}