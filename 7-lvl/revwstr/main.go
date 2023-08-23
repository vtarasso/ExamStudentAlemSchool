package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		return
	} else {
		res := ""
		arr := []string{}
		for _, w := range args[0] {
			if w != ' ' {
				res = res + string(w)
			} else if w == ' ' && res != "" {
				arr = append(arr, res)
				res = ""
			}
		}
		if res != "" {
			arr = append(arr, res)
		}
		for i := len(arr) - 1; i >= 0; i-- {
			for _, v := range arr[i] {
				z01.PrintRune(v)
			}
			if i != 0 {
				z01.PrintRune(' ')
			}
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
	// if len(os.Args) == 2 {
	var res []string
	word := ""
	for _, v := range os.Args[1] {
		if v != ' ' {
			word += string(v)
		}
		if v == ' ' {
			res = append(res, word)
			word = ""
		}
	}
	if word != "" {
		res = append(res, word)
	}
	for i := len(res) - 1; i >= 0; i-- {
		for _, j := range res[i] {
			z01.PrintRune(j)
		}
		if i != 0 {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}*/
