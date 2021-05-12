package main

import (
	"fmt"
	"reflect"
)

// 字段标签
func main() {
	type user struct {
		name string `昵称`
		sex  byte   `性别`
	}

	u := user{"Tom", 1}
	// 反射获取标签信息
	v := reflect.ValueOf(u)
	t := v.Type()

	for i, n := 0, t.NumField(); i < n; i++ {
		fmt.Printf("%s: %v\n", t.Field(i).Tag, v.Field(i))
	}
}
