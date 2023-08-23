package main
import (
	"os"
	"github.com/01-edu/z01"
)
func main() {
	if len(os.Args) != 2 {
		return
	}
	arg := os.Args[1]
	for _, w := range arg {
		for i := 0; i < Repeat(w); i++ {
			z01.PrintRune(w)
		}
	}
	z01.PrintRune('\n')
}
func Repeat(r rune) int {
	if r >= 'a' && r <= 'z' {
		return int(r - 96)
	} else if r >= 'A' && r <= 'Z' {
		return int(r - 64)
	} else {
		return 1
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
			if v >= 'a' && v <= 'z' {
				for i := 0; i <= int(v-'a'); i++ {
					z01.PrintRune(v)
				}
			} else if v >= 'A' && v <= 'Z' {
				for i := 0; i <= int(v-'A'); i++ {
					z01.PrintRune(v)
				}
			} else {
				z01.PrintRune(v)
			}
		}
		z01.PrintRune('\n')
	}
}
*/