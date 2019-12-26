package main

import (
	"fmt"
	"reflect"
)

/**
反射类型指的是reflect.Type和reflect.Value这两种
*/
func main() {
	var num float32 = 1.2345

	pointer := reflect.ValueOf(&num)
	value := reflect.ValueOf(num)

	fmt.Printf("num address :0x%0x , type:%s \n", &num, reflect.TypeOf(num))

	// 可以理解为“强制转换”，但是需要注意的时候，转换的时候，如果转换的类型不完全符合，则直接panic
	// Golang 对类型要求非常严格，类型一定要完全符合
	// 如下两个，一个是*float64，一个是float64，如果弄混，则会panic
	convertPointer := pointer.Interface().(*float64)
	convertValue := value.Interface().(float64)
	fmt.Println(convertPointer, *convertPointer)
	fmt.Println(convertValue)
}