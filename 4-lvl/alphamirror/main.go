package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		z01.PrintRune('\n') //return
	} else {
		pralph := []rune(args[0])
		for i, v := range pralph {
			if v >= 'a' && v <= 'z' {
				pralph[i] = 'a' + 'z' - v
			} else if v >= 'A' && v <= 'Z' {
				pralph[i] = 'A' + 'Z' - v
			}
		}
		for _, z := range string(pralph) {
			z01.PrintRune(z)
		}
		z01.PrintRune('\n')
	}
}

/*
package main
import (
	"os"
	"github.com/01-edu/z01"
)
func main() {
	if len(os.Args) == 2 {
		for _, v := range os.Args[1] {
			if v >= 'A' && v <= 'Z' {
				v = ('Z' - v + 'A')
			} else if v >= 'a' && v <= 'z' { // val >= 97 && val <= 122
				v = ('z' - v + 'a') // 122 - 98 + 97 =121
			}
			z01.PrintRune(v)
		}
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
	if len(os.Args) == 2 {
		for _, i := range os.Args[1] {
			if i >= 'A' && i <= 'Z' {
				i = ('Z' - (i - 'A'))
			} else if i >= 'a' && i <= 'z' {
				i = ('z' - (i - 'a'))
			}
			z01.PrintRune(i)
		}
		z01.PrintRune('\n')
	}
}
*/
