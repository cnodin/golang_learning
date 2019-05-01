package main

import "fmt"

type Data struct {
	complax []int //int型切片
	instance InnerData
	ptr *InnerData
}

type InnerData struct {
	a int
}

func passByValue(inFunc Data) Data {
	//输出参数的成员情况
	fmt.Printf("inFunc value: %+v\n", inFunc)

	//打印inFunc的指针
	fmt.Printf("inFunc ptr: %p\n", &inFunc)

	return inFunc
}

func main() {

	in := Data{
		complax: []int{1,2,3},
		instance: InnerData{
			5,
		},
		ptr: &InnerData{1},
	}
	fmt.Printf("in value: %+v\n", in)
	fmt.Printf("in ptr: %p\n", &in)

	out := passByValue(in)

	fmt.Printf("out value: %+v\n", out)
	fmt.Printf("out ptr: %p\n", &out)
}
