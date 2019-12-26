package main

import "fmt"

const (
	a = iota
	b
	c = 100
	d = iota		// 恢复自增长
	e
)

func main() {
	fmt.Println(a, b, c, d, e)
}
