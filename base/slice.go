package main

import (
	"fmt"
)

/**
slice 本身是个只读对象，其工作机制类似数组指针的一种包装
Edit by: jacky
*/
func main() {
	//x := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//p := &x
	//// cap 表示slice引用数组片段的真实长度 cap = max-low = 7-2 = 5
	//a := x[2:5:7]
	//printSlice(a)
	//
	//p[2] += 10
	//printSlice(a)
	//
	//// 直接初始化 slice
	//fmt.Println("\n====slice directly===")
	//s1 := make([]int, 3, 5)
	//s2 := make([]int, 3)
	//s3 := []int{10, 20, 5: 30}
	//printSlice(s1)
	//printSlice(s2)
	//printSlice(s3)
	//
	//// diff
	//var x1 []int
	//y1 := []int{} //initialed
	//fmt.Println("x1 type = ", reflect.TypeOf(x1))
	//
	//fmt.Println(x1 == nil, y1 == nil, unsafe.Sizeof(x1))
	//fmt.Printf("x1: %#v\n", (*reflect.SliceHeader)(unsafe.Pointer(&x1)))
	//fmt.Printf("y1: %#v\n", (*reflect.SliceHeader)(unsafe.Pointer(&y1)))

	append_slice()
}

//func printSlice(a []int) {
//	fmt.Printf("len=%d,cap=%d\t", len(a), cap(a))
//	for i := 0; i < len(a); i++ {
//		fmt.Printf("a[%d]=%d\t", i, a[i])
//	}
//	fmt.Println()
//}


func append_slice() {
	var a []int
	printSlice("a", a)

	// append works on nil slices.
	a = append(a, 0)
	printSlice("a", a)

	// the slice grows as needed.
	a = append(a, 1)
	printSlice("a", a)

	// we can add more than one element at a time.
	a = append(a, 2, 3, 4)
	printSlice("a", a)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
