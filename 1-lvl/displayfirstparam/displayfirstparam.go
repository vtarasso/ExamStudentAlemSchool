package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	if len(arg) < 2 {
			return
		}
	for _, args := range arg[0] {
		z01.PrintRune(args)
	}
	z01.PrintRune('\n')
}

/*
displaylastparam
	func main() {
		arg := os.Args
		args := arg[len(arg)-1]
		if len(arg) < 2 {
			return
		}
		for _, i := range args {
			z01.PrintRune(i)
		}
		z01.PrintRune('\n')
	}
*/
