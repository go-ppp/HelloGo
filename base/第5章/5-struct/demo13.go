package main

import (
	"fmt"
	"unsafe"
)

// 内存布局，孔结构类型字段，在最后一个字段
func main() {
	v:= struct {
		a struct{}
		b int
		c struct{}
	}{}

	s := `
		v: %p ~ %x ,size: %d ,align: %d

		field 	address 		offset 	size
		------+---------------+--------+------------
		a		%p				%d		%d
		b		%p				%d		%d
		c		%p				%d		%d
		`

	fmt.Printf(s,
		&v, uintptr(unsafe.Pointer(&v))+unsafe.Sizeof(v), unsafe.Sizeof(v), unsafe.Alignof(v),
		&v.a, unsafe.Offsetof(v.a), unsafe.Sizeof(v.a),
		&v.b, unsafe.Offsetof(v.b), unsafe.Sizeof(v.b),
		&v.c, unsafe.Offsetof(v.c), unsafe.Sizeof(v.c))
}
