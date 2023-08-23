package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		res := ""
		arr := []string{}
		for _, w := range os.Args[1] {
			if w != ' ' {
				res = res + string(w)
			} else if res != "" && w == ' ' {
				arr = append(arr, res)
				res = ""
			}
		}
		if res != "" {
			arr = append(arr, res)
		}
		str := ""
		for _, v := range arr[1:] {
			str = str + v + " "
		}
		str1 := str + arr[0]
		for _, z := range str1 {
			z01.PrintRune(z)
		}
		z01.PrintRune('\n')
	}
}

/*package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	var res []string
	word := ""
	for _, v := range os.Args[1] {
		if v != ' ' {
			word += string(v)
		} else if word != "" {
			res = append(res, word)
			word = ""
		} else {
			continue
		}
	}
	if word != "" {
		res = append(res, word)
	}
	val := ""
	for i := len(res) - 1; i > 0; i-- {
		val = res[i] + " " + val
	}
	val += res[0]
	for _, v := range val {
		z01.PrintRune(v)
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
	args := os.Args[1:]
	if len(args) == 1 {
		var res []string
		word := ""
		for _, v := range os.Args[1] {
			if v != ' ' {
				word = word + string(v)
			} else if v == ' ' && word != "" {
				res = append(res, word)
				word = ""
			} else {
				continue
			}
		}
		if word != "" {
			res = append(res, word)
		}
		str := ""
		for _, w := range res[1:] {
			str = str + w + " "
		}
		str1 := str + res[0]
		for _, n := range str1 {
			z01.PrintRune(n)
		}
		z01.PrintRune('\n')
	} else {
		z01.PrintRune('\n')
		return
	}
}
*/
