package main

// 别名
func test(x byte) {
	println(x)
}

func main() {
	var a byte = 0x11
	var b uint8 = a
	var c uint8 = a + b

	test(c)
}
