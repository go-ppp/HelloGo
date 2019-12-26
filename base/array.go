package main

import "fmt"

type user struct {
	name string
	age  int8
}

func testP(x [2]int) {
	fmt.Printf("x: %p, %v\n", &x, x)
}

func testP2(x *[2]int) {
	fmt.Printf("x: %p, %v\n", x, *x)
}

/**
复制： go数组是值类型， 赋值和传参都会复制整个数组数据
用指针或slice，避免数据复制
*/
func main() {
	d1 := [3]int{1, 2}
	for i, c := range d1 {
		fmt.Printf("d1[%d] = %d\t", i, c)
	}

	fmt.Println()
	d2 := [4]int{5, 2: 10} // 索引方式初始化
	for i, c := range d2 {
		fmt.Printf("d2[%d] = %d\t", i, c)
	}

	d3 := [...]user{
		{"Tom", 20},
		{"Jacky", 30},
	}
	fmt.Printf("\nusers-> %#v\n", d3)

	// 指针数组 和 数组指针
	a := [...]int{1, 2}
	p := &a
	p[1] += 10
	fmt.Println(a, &a, &a[0], &a[1])

	// 复制
	var b [2]int
	b = a
	fmt.Printf("a: %p, %v\n", &a, a)
	fmt.Printf("b: %p, %v\n", &b, b)
	testP(a)

	// 用指针或slice，避免数据复制
	fmt.Printf("a: %p, %v\n", &a, a)
	testP2(&a)
}
