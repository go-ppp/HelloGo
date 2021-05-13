package main

import "fmt"

// 浮点数多精度,可以参考 IEEE-754
func main() {
	var a float32 = 1.1234567899
	var b float32 = 1.12345678
	var c float32 = 1.123456781

	println(a, b, c)

	println(a == b, a == c)
	fmt.Printf("%v %v, %v\n", a, b, c)
}
