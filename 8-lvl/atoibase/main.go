package main

import (
	"fmt"
)

func main() {
	fmt.Println(AtoiBase("125", "0123456789"))
	fmt.Println(AtoiBase("1111101", "01"))
	fmt.Println(AtoiBase("7D", "0123456789ABCDEF"))
	fmt.Println(AtoiBase("uoi", "choumi"))
	fmt.Println(AtoiBase("bbbbbab", "-ab"))
}

func AtoiBase(s string, base string) int {
	m := make(map[rune]int)
	for i, w := range base {
		if w == '+' || w == '-' {
			return 0
		}
		m[w] = i
	}
	var n int
	for _, v := range s {
		n = n*len(base) + m[v]
	}
	return n
}
