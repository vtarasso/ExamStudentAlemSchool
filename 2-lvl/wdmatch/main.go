package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	if len(arg) != 2 {
		return
	}
	for i, j := 0, 0; i < len(arg[0]) && j < len(arg[1]); {
		if arg[0][i] == arg[1][j] {
			i++
			if i == len(arg[0]) {
				for _, w := range arg[0] {
					z01.PrintRune(w)
				}
				z01.PrintRune('\n')
			}
		}
		j++
		if j == len(arg[1]) {
			return
		}
	}
}

/*
package main

import (
	"os"


	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	if len(args) != 3 {
		return
	}
	a := os.Args[1]
	b := os.Args[2]
	c := 0
	var s string
	for i := 0; i < len(a); i++ {
		for j := c + 1; j < len(b); j++ {
			if a[i] == b[j] {
				s += string(a[i])
				j = len(b) - 1
			}
			c++
		}
	}
	if s == a {
		for _, k := range s {
			z01.PrintRune(k)
		}
		z01.PrintRune('\n')
	}
}

*/

/*
	func main() {
	args := os.Args
	if len(args) != 3 {
		return
	}
	s1 := args[1]
	s2 := args[2]
	count := 0
	j := 0
	for i := range s2 {
		if j >= len(s1) {
			break
		}
		if s2[i] == s1[j] {
			j++
			count++
		}
	}
	if count == len(s1) {
		for _, k := range s1 {
			z01.PrintRune(k)
		}
		z01.PrintRune('\n')
	} else {
		return
	}
}
*/
