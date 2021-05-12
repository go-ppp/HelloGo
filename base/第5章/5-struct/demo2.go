package main

import "fmt"

func main() {
	type user struct {
		name string
		age  byte
	}

	u1 := user{"Tom", 12}
	u2 := user{"Tom"} // wrong

	fmt.Println(u1, u2)
}
