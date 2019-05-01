package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Millisecond * 50)

	stopper := time.NewTimer(time.Second * 2)

	var i int

	for {
		select {
		case <-stopper.C:
			fmt.Println("stop")
			goto StopHere
		case <-ticker.C:
			i++ //记录触发了多少次
			fmt.Println("tick", i)
		}
	}

StopHere:
	fmt.Println("finished")
}
