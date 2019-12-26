package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	a, b, c := 100, 0144, 0x64
	fmt.Println(a, b, c)
	fmt.Printf("0b%b, %#o, %#x\n", a, b, c)
	fmt.Println(math.MinInt8, math.MaxInt8)

	stringConvert()
}

func stringConvert() {
	const s = "1011010"
	a, _ := strconv.ParseInt(s, 2, 32)
	fmt.Println(s, a)

	println("0b" + strconv.FormatInt(a, 2))
	println("0" + strconv.FormatInt(a, 8))
	println("0x" + strconv.FormatInt(a, 16))
}
