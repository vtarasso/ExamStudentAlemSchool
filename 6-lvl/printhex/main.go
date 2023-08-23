package main

import (
	"os"
	"strconv"
	"github.com/01-edu/z01"
)

func Print(s string) {
	for _, str := range s {
		z01.PrintRune(str)
	}
	z01.PrintRune('\n')
}
func main() {
	if len(os.Args) != 2 {
		return
	}
	arg, err := strconv.Atoi(os.Args[1])
	if err != nil {
		Print("ERROR")
		return
	}
	hex := strconv.FormatInt(int64(arg), 16)
	Print(hex)
}
/*package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	ar := os.Args
	errStr := "ERROR"
	if len(ar) == 2 {
		n, er := strconv.Atoi(ar[1])
		if er != nil {
			for _, v := range errStr {
				z01.PrintRune(v)
			}
		}
		for _, r := range hex(n) {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}

func hex(n int) string {
	base := "0123456789abcdef"
	if n < 16 {
		return string(base[n])
	}
	return hex(n/16) + string(base[n%16])
}
*/

/*
package main
import (
	"os"
	"strconv"
	"github.com/01-edu/z01"
)
func main() {
	if len(os.Args) != 2 {
		return
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		for _, v := range "ERROR" {
			z01.PrintRune(v)
		}
		return
	}
	res := strconv.FormatInt(int64(n), 16)
	for _, v := range res {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}
*/