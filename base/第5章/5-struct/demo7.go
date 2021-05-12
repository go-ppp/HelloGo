package main

import (
	"fmt"
)

// 匿名字段
func main() {

	type attr struct {
		perm int
	}

	type file struct {
		name string
		attr // 匿名字段
	}

	f := file{
		name: "test.dat",
		attr: attr{
			perm: 0755,
		},
	}

	fmt.Println(f)
	f.perm = 0644
	fmt.Println(f)
}
