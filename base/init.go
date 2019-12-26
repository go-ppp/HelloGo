package main

import "fmt"

func main() {
	type data struct {
		name string
		age  int8
	}

	var a data = data{"jacky", 20}

	b := data{
		"alice",
		100,
	}

	fmt.Println(a, b)
}
