package main

import "fmt"

type Data struct {
	complax []int
	instance InnerData
	ptr *InnerData
}

type InnerData struct {
	a int
}

func passByValue(inFunc Data) Data {
	fmt.Printf("inFunc value: %+v\n", inFunc)

	fmt.Printf("inFunc ptr: %p\n", &inFunc)

	return inFunc
}

func fire() {
	fmt.Println("fire")
}

func main() {

	in := Data{
		complax: []int{1, 2 , 3},
		instance: InnerData {
			5,
		},

		ptr: &InnerData{1},
	}

	//输入结构体的成员情况
	fmt.Printf("in value: %+v\n", in)

	//输入结构体的指针地址
	fmt.Printf("in ptr: %p\n", &in)

	out := passByValue(in)

	//输出结构的成员情况
	fmt.Printf("out value: %+v\n", out)

	//输出结构的指针地址
	fmt.Printf("out ptr: %p\n", &out)

	var f func()
	f = fire
	f()
}
