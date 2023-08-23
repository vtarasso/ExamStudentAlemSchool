package main

import "fmt"

func Max(a []int) int {
	var max int
	for i := range a {
		if max < a[i] {
			max = a[i]
		}
	}
	return max
}

func main() {
	a := []int{23, 123, 1, 11, 55, 93}
	max := Max(a)
	fmt.Println(max)
}
