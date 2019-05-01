package main

func a() {}
func b() {}

//定义函数类型
//type FormatFunc func(string, ...interface{})(string, error)
//
//func format(f FormatFunc, s string, a ...interface{}) (string, error)  {
//	return f(s, a)
//}

func test() *int {
	a := 0x100
	return &a
}

func main() {
	//println(a == nil)
	//println(a == b) //编译不通过，函数只能和nil比较

	var a *int = test()
	println(a, *a)
}
