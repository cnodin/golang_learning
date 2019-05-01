package main

import (
	"reflect"
	"fmt"
)

//商标结构
type Brand struct {

}

//为商标结构添加show方法
func (t Brand) Show() {

}

type FakeBrand = Brand

type Vehicle struct {
	FakeBrand
	Brand
}

func main() {

	var a Vehicle

	a.FakeBrand.Show()

	//取a的类型反射对象
	ta := reflect.TypeOf(a)

	//遍历a的所有成员
	for i := 0; i < ta.NumField(); i++ {
		f := ta.Field(i)

		fmt.Printf("FieldName: %v, FieldType: %v\n", f.Name, f.Type.Name())
	}

}
