package main

import (
	"errors"
	"log"
)

func main() {
	a, e := div(10, 0)
	if e != nil {
		log.Fatalln(e)
	}

	println(a)
}

var errDivByZero = errors.New("division by zero")

func div(x, y int) (int, error) {
	if y == 0 {
		return 0, errDivByZero
	}
	return x / y, nil
}
