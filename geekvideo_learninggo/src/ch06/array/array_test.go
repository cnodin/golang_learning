package array_test

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr1 := [3]int{1,2,3}
	arr2 := [...]int{1,2,3,4}
	t.Log(arr[1], arr[2])
	t.Log(arr1, arr2)
}

func TestArrayTravel(t *testing.T) {
	arr1 := [...]int{1,2,3,4}
	for i:= 0; i < len(arr1); i++ {
		t.Log(arr1[i])
	}
	for _, e := range arr1 {
		t.Log(e)
	}

	//a[开始索引(包含)，结束索引(不包含)]
	t.Log(arr1[1:2])
	t.Log(arr1[1:3])
	t.Log(arr1[0:len(arr1)])
}
