package main

import (
	"runtime"
	"sync"
	"fmt"
)

func main() {

	//Allocate 1 logic process for the scheduler to user
	runtime.GOMAXPROCS(runtime.NumCPU())

	//wg is used to wait for the program to finish.
	//Add a count of two, one for each goroutine
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutines")

	//declare an anonymous function and create a goroutine
	go func() {
		defer wg.Done()

		//display the alphabet three times
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a' + 26; char++ {
				fmt.Printf("%c  ", char)
			}
		}
	}()

	//declare an anonymous function and create a goroutine
	go func() {
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A' + 26; char++ {
				fmt.Printf("%c  ", char)
			}
		}
	}()

	//wait for the goroutines to finish
	fmt.Println("Waiting to Finish")
	wg.Wait()

	fmt.Println("\nTerminating Program")
}
