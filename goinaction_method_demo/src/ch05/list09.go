package main

import (
	"sync"
	"runtime"
	"fmt"
	"sync/atomic"
)

var (
	counter int64
	wg sync.WaitGroup
)

func main() {
	wg.Add(2)

	go intCounter(1)
	go intCounter(2)

	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func intCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2 ; count++  {
		atomic.AddInt64(&counter, 1)
		runtime.Gosched()
	}
}
