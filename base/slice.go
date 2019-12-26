package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

/**
slice 本身是个只读对象，其工作机制类似数组指针的一种包装
*/
func main() {
	x := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	p := &x
	// cap 表示slice引用数组片段的真实长度 cap = max-low = 7-2 = 5
	a := x[2:5:7]
	printSlice(a)

	p[2] += 10
	printSlice(a)

	// 直接初始化 slice
	fmt.Println("\n====slice directly===")
	s1 := make([]int, 3, 5)
	s2 := make([]int, 3)
	s3 := []int{10, 20, 5: 30}
	printSlice(s1)
	printSlice(s2)
	printSlice(s3)

	// diff
	var x1 []int
	y1 := []int{} //initialed
	fmt.Println(x1 == nil, y1 == nil, unsafe.Sizeof(x1))
	fmt.Printf("x1: %#v\n", (*reflect.SliceHeader)(unsafe.Pointer(&x1)))
	fmt.Printf("y1: %#v\n", (*reflect.SliceHeader)(unsafe.Pointer(&y1)))
}

func printSlice(a []int) {
	fmt.Printf("len=%d,cap=%d\t", len(a), cap(a))
	for i := 0; i < len(a); i++ {
		fmt.Printf("a[%d]=%d\t", i, a[i])
	}
	fmt.Println()
}
