package main

import (
	"fmt"
	"math"
)

func add(base int) func(int) int {
	return func(i int) int {
		base += i
		return base
	}
}

func test01(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}
	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func main() {
	tmp1 := add(10)
	tmp2 := add(100)
	fmt.Println(tmp1(1), tmp1(2))
	fmt.Println(tmp2(1), tmp2(2))
	f1, f2 := test01(10)
	fmt.Println(f1(1), f2(2))
	fmt.Println(f1(3), f2(4))

	getSqrt := func(a float64) float64 {
		return math.Sqrt(a)
	}
	fmt.Println(getSqrt(6.1))
}
