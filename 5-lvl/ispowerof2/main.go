package main
import (
	"os"
	"strconv"
	"github.com/01-edu/z01"
)
func isPowerOf2(n int) bool {
	return n != 0 && ((n & (n - 1)) == 0)
}
func main() {
	if len(os.Args) == 2 {
		nbr, _ := strconv.Atoi(os.Args[1])
		if isPowerOf2(nbr) {
			for _, ch := range "true" {
				z01.PrintRune(ch)
			}
		} else {
			for _, ch := range "false" {
				z01.PrintRune(ch)
			}
		}
		z01.PrintRune('\n')
	}
}

/*
package main
import (
	"os"
	"strconv"
	"github.com/01-edu/z01"
)
func main() {
	if len(os.Args) == 2 {
		n, ea := strconv.Atoi(os.Args[1])
		if ea != nil {
			z01.PrintRune('\n')
			return
		}
		if isPowerofTwo(n) {
			for _, r := range "true" {
				z01.PrintRune(r)
			}
		} else {
			for _, r := range "false" {
				z01.PrintRune(r)
			}
		}
	}
	z01.PrintRune('\n')
}
func isPowerofTwo(n int) bool {
	if n&(n-1) == 0 && n != 0 { // битовая операция //1000 - 8 в десятичной системе
		// 0111 - 7 в десятичной системе
		return true
	}
	return false
}
*/

/*
package main
import (
	"os"
	"strconv"
	"github.com/01-edu/z01"
)
//"strconv"
func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		return
	}
	n, _ := strconv.Atoi(args[0])
	// if err == nil {
	// for i := 0; i < len(args[0]); i++ {
	if n&(n-1) == 0 && n > 0 {
		res := "true"
		for _, v := range res {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')
	} else {
		res := "false"
		for _, v := range res {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')
	}
}
func Atoi(args string) int {
	s := 0
	temp := 1
	if args[0] == '-' {
		temp = -1
		args = args[1:]
	}
	for v := range args {
		s = s * 10
		s = s + int(args[v]-48)
	}
	return s * temp
}
*/