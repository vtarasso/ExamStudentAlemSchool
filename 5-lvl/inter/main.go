package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		return
	} else {
		res := ""
		for _, w := range args[0] {
			for _, v := range args[1] {
				if w == v {
					if !Choice(res, w) && w != ' ' {
						res = res + string(w)
					}
				}
			}
		}
		for _, w := range res {
			z01.PrintRune(w)
		}
		z01.PrintRune('\n')
	}
}

func Choice(res string, w rune) bool {
	for _, z := range res {
		if z == w {
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
			for _, b := range s2 {
				if a == b {
					if _, value := check[string(a)]; !value {
						check[string(a)] = true
						z01.PrintRune(a)
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}*/
/*
package main
import (
	"os"
	"github.com/01-edu/z01"
)
func main() {
	if len(os.Args) != 3 {
		return
	}
	s1 := os.Args[1]
	s2 := os.Args[2]
	res := ""
	for _, v1 := range s1 {
		for _, v2 := range s2 {
			if v1 == v2 && Contain(res, v1) {
				res += string(v1)
			}
		}
	}
	for _, v := range res {
		z01.PrintRune(v)
	}
	z01.PrintRune(10)
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
	} else {
		res := ""
		for _, w := range args[0] {
			for _, v := range args[1] {
				if w == v {
					if !Choice(res, w) && w != ' ' {
						res = res + string(w)
					}
				}
			}
		}
		for _, w := range res {
			z01.PrintRune(w)
		}
		z01.PrintRune('\n')
	}
}
func Choice(res string, w rune) bool {
	for _, z := range res {
		if z == w {
			return true
		}
	}
	return false
}
*/
