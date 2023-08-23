package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 3 {
		res := ""
		temp := 0
		for i := 0; i < len(os.Args[1]); i++ {
			for j := temp; j < len(os.Args[2]); j++ {
				if os.Args[1][i] == os.Args[2][j] {
					res = res + string(os.Args[1][i])
					temp = j + 1
					break
				}
			}
		}
		if len(res) == len(os.Args[1]) {
			z01.PrintRune('1')
		} else {
			z01.PrintRune('0')
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
	if len(os.Args) != 3 {
		return
	}
	defer z01.PrintRune('\n')
	count := 0
	for _, v := range os.Args[2] {
		if len(os.Args[1]) == count {
			z01.PrintRune('1')
			return
		}
		if v == rune(os.Args[1][count]) {
			count++
		}
	}
	z01.PrintRune('0')
}
*/
/*package main
import (
	"fmt"
	"os"
)
func main() {
	if len(os.Args) == 3 {
		s1 := os.Args[1]
		s2 := os.Args[2]
		i := 0
		res := ""
		for _, v := range s2 {
			if len(res) == len(s1) {
				break
			}
			if v == rune(s1[i]) {
				res += string(v)
				i++
			}
		}
		if res == s1 {
			fmt.Println(1)
			return
		}
		fmt.Println(0)
	}
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
		temp := 0
		for i := 0; i < len(args[0]); i++ {
			for j := temp; j < len(args[1]); j++ {
				if args[0][i] == args[1][j] {
					temp = j + 1
					res = res + string(args[0][i])
					break
				}
			}
		}
		if len(res) == len(args[0]) {
			z01.PrintRune('1')
		} else {
			z01.PrintRune('0')
		}
		z01.PrintRune('\n')
	}
}
*/
