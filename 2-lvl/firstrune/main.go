package main

import "github.com/01-edu/z01"

func FirstRune(s string) rune {
	a := []rune(s)
	return a[0]
	// return rune(s[0])
}

func main() {
	z01.PrintRune(FirstRune("Salut!"))
	z01.PrintRune(FirstRune("Hello!"))
	z01.PrintRune(FirstRune("Ola!"))
	z01.PrintRune('\n')
}

/*
lastRune
	func LastRune(s string) rune {
		//return rune(s[len(s)-1])
		a := []rune(s)
		return a[len(s)-1]
	}
*/