package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	res := strconv.Itoa(len(os.Args[1:]))
	for _, w := range res {
		z01.PrintRune(w)
	}
	z01.PrintRune('\n')
}

/*package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	zprint(len(os.Args[1:]))
}

func Print(n int) {
	runes := []rune(strconv.Itoa(n))
	for _, v := range runes {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}
*/
/*
package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	arg := len(args)

	runes := []rune{}

	for i := arg; i > 0; i = i / 10 {
		runes = append(runes, rune(i%10+48))
	}
	for v := len(runes) - 1; v >= 0; v-- {
		z01.PrintRune(runes[v])
	}
	if arg == 0 {
		z01.PrintRune(rune(48))
	}
	z01.PrintRune('\n')
}
*/
/*
	func main() {
		arg := os.Args[1:]
		len := len(arg)
		if len > 9 {
			num1 := len / 10
			num2 := len % 10
			z01.PrintRune(rune(num1) + 48)
			z01.PrintRune(rune(num2) + 48)
			z01.PrintRune('\n')
		} else {
			z01.PrintRune(rune(len + '0'))
			z01.PrintRune('\n')
		}
	}
*/
