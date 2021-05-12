package main

import "fmt"

func main() {
	a := [10]struct{}{}
	b := a[:]
	c := [0]int{}

	// 长度为0的元素，通常都指向 runtime.zerobase 变量
	fmt.Printf("%p, %p, %p\n", &a[0], &b[0], &c)
}
