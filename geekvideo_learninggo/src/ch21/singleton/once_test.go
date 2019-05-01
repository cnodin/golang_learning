package once_test

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

type Singleton struct {
	a bool
}

var singleInstance *Singleton

var once sync.Once

func GetSingletonObj() *Singleton {
	//once.Do(func() {
		fmt.Println("create obj")
		singleInstance = new(Singleton)
		singleInstance.a = false
	//})

	return singleInstance
}

func TestGetSingletonObj(t *testing.T)  {
	var wg sync.WaitGroup
	for i:=0; i < 10; i++ {
		wg.Add(1)
		go func() {
			obj := GetSingletonObj()
			fmt.Printf("%x\n", unsafe.Pointer(obj))
			//fmt.Println(obj)
			wg.Done()
		}()
	}
	wg.Wait()
}