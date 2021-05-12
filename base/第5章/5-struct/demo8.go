package main

import (
	"fmt"
	"os"
)

// 匿名字段
func main() {
	type data struct {
		os.File
	}

	d := data{
		File: os.File{},
	}

	fmt.Printf("%#v\n", d)
}
