package share_mem_test

import (
	"sync"
	"testing"
	"time"
)

func TestCounter(t *testing.T)  {
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 50000000; i++ {
		go func() {
			defer func() {
				mut.Unlock()
			}()

			mut.Lock()
			counter++
		}()
	}
	time.Sleep(time.Microsecond * 50)
	t.Logf("Counter = %d", counter)
}

func TestWaitGroup (t *testing.T)  {
	var mut sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				mut.Unlock()
			}()

			mut.Lock()
			counter++
			wg.Done()
		}()
	}

	wg.Wait()
	t.Logf("counter = %d", counter)
}
