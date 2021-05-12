package main

import "fmt"

func main() {
	// 匿名结构类型变量
	u := struct {
		name string
		age  byte
	}{
		name: "Tom",
		age:  12,
	}

	// 定义匿名结构类型字段
	type file struct {
		name string
		attr struct {
			owner int
			perm  int
		}
	}

	f := file{
		name: "Tom",
		//attr:  {				// wrong
		//	owner :1,
		//	perm : 0.775
		//},
	}

	f.attr.owner = 1
	f.attr.perm = 0775

	fmt.Println(u, f)
}
