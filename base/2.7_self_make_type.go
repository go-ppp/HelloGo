package main

import "fmt"

// 2.7  自定义类型
func main() {
	type (
		user struct {
			name string
			age  uint8
		}

		event func(string) bool // 函数类型
	)

	u := user{"Jacky", 20}
	fmt.Println(u)

	var f event = func(s string) bool {
		println(s)
		return s != ""
	}

	const s = "abc"
	if f(s) {
		fmt.Println(s + " is not empty")
	}
}
