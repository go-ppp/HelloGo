package main

import "fmt"

// 指针
/**
1. & 用于获取对象地址
2. "*" 用于间接引用目标对象

指针类型不能做加减法
*/
func main() {
	// first
	//x := 10
	//var p *int = &x
	//*p += 20
	//fmt.Println(p, *p, x)

	// second, 指针没有转门指向成员的 "->" 运算符，统一用"."选择表达式
	a := struct {
		x int
	}{}

	a.x = 100
	p := &a
	p.x += 100
	fmt.Println(p, &p, p.x)
}
