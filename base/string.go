package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

/**
要修改字符串，必须将其转换为可变类型（[]rune 或者 []byte）待完成后在转换回来。 但不管怎么转换都需要重新分配内存，并复制数据
*/
func main() {
	s := "abc1234"
	p := &s
	println(*p, len(s))

	for i := 0; i < len(s); i++ { // byte 方式
		if i == 0 {
			fmt.Println("byte 方式遍历:")
		}
		fmt.Print(s[i], "\t")

		if i == len(s)-1 {
			fmt.Println()
		}
	}

	for i, c := range s { // rune 方式
		if i == 0 {
			fmt.Println("rune 方式遍历")
		}
		fmt.Printf("s[%d] = %d\t", i, c)

		if i == len(s)-1 {
			fmt.Println()
		}
	}

	// slice
	s1 := s[:3]
	s2 := s[1:4]
	s3 := s[2:]
	println("slice: ", s1, s2, s3)

	convert_test()
}

func pp(format string, ptr interface{}) {
	p := reflect.ValueOf(ptr).Pointer()
	h := (*uintptr)(unsafe.Pointer(p))
	fmt.Printf(format, *h)
}

func convert_test() {
	s := "hello world"
	pp("s: %x\n", &s)

	bs := []byte(s)
	s2 := string(bs)

	pp("string to []byte, bs: %x\n", &bs)
	pp("[]byte to string, s2: %x\n", &s2)
}
