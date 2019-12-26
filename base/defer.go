package main

/**
defer:

3. 多个延迟注册按FiILO次序进行
4. return 和 panic 语句会终止当前函数流程，引发延迟调用. 注意： 延迟调用是在函数结束时才会执行
5. 性能比较，vs call
*/
func main() {
	// 1
	//f, err := os.Open("./base_data.go")
	//if (err != nil) {
	//	log.Fatalln(err)
	//} else {
	//	fmt.Println(f.Name())
	//}
	//
	//defer f.Close()
	//
	//fmt.Println("do something")

	// 2
	//x, y := 1, 2
	//defer func(a int) {
	//	println("defer x, y = ", a, y)		// y为闭包引用
	//}(x)					// 注册时复制调用参数
	//
	//x += 100
	//y += 200
	//println(x, y)

	// 3
	//defer println("a")
	//defer println("b")

	// 4
	//println("test:", test())
}

func test() (z int) {
	defer func() {
		println("defer:", z)
		z += 100
	}()

	return 100
}
