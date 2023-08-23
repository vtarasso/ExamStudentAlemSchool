package main

import "github.com/01-edu/z01"

func main() {
	ch2 := 'A'
	for ch1 := 'a'; ch1 <= 'z'; ch1++ {
		if ch1%2 == 1 {
			z01.PrintRune(ch1)
		} else {
			z01.PrintRune(ch2)
		}
		ch2++
	}
	z01.PrintRune('\n')
}

/*
displayalrevm
	func main() {
		ch2 := 'z'
		for ch1 := 'Z'; ch1 >= 'A'; ch1-- {
			if ch1%2 == 1 {
				z01.PrintRune(ch1)
			} else {
				z01.PrintRune(ch2)
			}
			ch2--
		}
		z01.PrintRune('\n')
	}
*/
