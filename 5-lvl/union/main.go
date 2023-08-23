package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		return
	}
	res := ""
	s := args[0] + args[1]
	for _, w := range s {
		if !Choice(res, w) {
			res = res + string(w)
		}
	}
	for _, v := range res {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}

func Choice(res string, w rune) bool {
	for _, n := range res {
		if n == w {
			return true
		}
	}
	return false
}

/*package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 3 {
		s1 := os.Args[1]
		s2 := os.Args[2]
		check := make(map[string]bool)
		for _, a := range s1 {
			if _, value := check[string(a)]; !value {
				check[string(a)] = true
				z01.PrintRune(a)
			}
		}
		for _, b := range s2 {
			if _, value := check[string(b)]; !value {
				check[string(b)] = true
				z01.PrintRune(b)
			}
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
	res := ""
	if len(os.Args) == 3 {
		str1 := os.Args[1]
		str2 := os.Args[2]
		for _, v := range str1 {
			if Contain(res, v) {
				res += string(v)
			}
		}
		for _, v := range str2 {
			if Contain(res, v) {
				res += string(v)
			}
		}
		for _, v := range res {
			z01.PrintRune(v)
		}
		z01.PrintRune(10)
	} else {
		z01.PrintRune(10)
	}
}
func Contain(s string, r rune) bool {
	for _, v := range s {
		if v == r {
			return false
		}
	}
	return true
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
	if len(args) != 2 {
		return
	}
	res := ""
	s := args[0] + args[1]
	for _, w := range s {
		if !Choice(res, w) {
			res = res + string(w)
		}
	}
	for _, v := range res {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}
func Choice(res string, w rune) bool {
	for _, n := range res {
		if n == w {
			return true
		}
	}
	return false
}
*/
