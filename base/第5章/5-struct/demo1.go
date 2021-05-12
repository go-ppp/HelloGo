package main

import "fmt"

type node struct {
	_    int
	id   int
	next *node
}

func main() {
	n1 := node{
		id: 1,
	}

	n2 := node{
		id:   2,
		next: &n1,
	}

	p1 := &n1
	fmt.Println(p1, n1, n2)
	// print address
	fmt.Printf("%p", p1)
}
