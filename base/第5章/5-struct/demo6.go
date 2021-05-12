package main

import (
	"fmt"
	"time"
)

// 空结构可以作为通道元素类型，用于事件通知
func main() {
	exit := make(chan struct{})

	go func() { // 协程
		time.Sleep(time.Second * 3)
		println("hello world")
		exit <- struct{}{}
	}()

	fmt.Println(1)
	<- exit
	//select {
	//case <-exit:
	//	fmt.Println("get value")
	//}
	fmt.Println("end.")
}
