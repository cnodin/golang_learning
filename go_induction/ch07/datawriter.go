package main

import "fmt"

type DataWriter interface {
	WriteData(data interface{}) error

	CanWrite() bool
}

type file struct {

}

func (f *file) WriteData(data interface{}) error {
	fmt.Println("WriteData:", data)
	return nil
}

func main() {

	f := new(file)

	var writer DataWriter

	writer = f

	writer.WriteData("data")


}
