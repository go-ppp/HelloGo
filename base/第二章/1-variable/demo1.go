package main

func main() {
	x := 100
	println(&x)

	x, y := 200, "abc"			// x,退化为赋值操作， 仅有 y 是变量定义

	println(&x, x)
	println(y)
}
