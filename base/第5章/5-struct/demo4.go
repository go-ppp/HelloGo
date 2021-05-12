package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a struct{}
	var b [100]struct{}

	println(unsafe.Sizeof(a), unsafe.Sizeof(b))

	var d [100]struct{}
	s := d[:]

	d[1] = struct{}{}
	d[2] = struct{}{}

	fmt.Println(s[3], len(s), cap(s))
}
