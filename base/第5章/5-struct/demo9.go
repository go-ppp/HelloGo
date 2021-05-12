package main

import "fmt"

// 匿名字段
func main() {
	type data struct {
		*int // 嵌入指针类型
		string
	}

	x := 100
	d := data{
		int:    &x, // 使用基础类型作为字段名
		string: "abc",
	}

	fmt.Printf("%#v\n", d)
}
