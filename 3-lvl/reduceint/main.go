package main

import (
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	mul := func(acc int, cur int) int {
		return acc * cur
	}
	sum := func(acc int, cur int) int {
		return acc + cur
	}
	div := func(acc int, cur int) int {
		return acc / cur
	}
	as := []int{500, 2}
	ReduceInt(as, mul)
	ReduceInt(as, sum)
	ReduceInt(as, div)
}

func ReduceInt(a []int, f func(int, int) int) {
	n := a[0]
	for i := 1; i < len(a); i++ {
		n = f(n, a[i])
	}
	arr := strconv.Itoa(n)
	for _, v := range arr {
		z01.PrintRune(v)
	}
	z01.PrintRune(10)
}

// func ReduceInt(a []int, f func(int, int) int) {
// 	res := f(a[0], a[1])
// 	z01.PrintRune(rune(res))
// }
