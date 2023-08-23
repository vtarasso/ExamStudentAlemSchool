package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		for _, ch := range os.Args[1] {
			if ch >= 'A' && ch <= 'Z' {
				if ch >= 'A'+13 {
					ch -= 13
				} else {
					ch += 13
				}
			}
			if ch >= 'a' && ch <= 'z' {
				if ch >= 'a'+13 {
					ch -= 13
				} else {
					ch += 13
				}
			}
			z01.PrintRune(ch)
		}
		z01.PrintRune('\n')
	}
}

/*

func main() {
	args := os.Args[1]
	var str string
	r := []rune(args)
	if len(os.Args) == 2 {
		for i := 0; i < len(r); i++ {
			if r[i] >= 'A' && r[i] <= 'M' || r[i] >= 'a' && r[i] <= 'm' {
				r[i] = r[i] + 13
			} else if r[i] >= 'N' && r[i] <= 'Z' || r[i] >= 'n' && r[i] <= 'z' {
				r[i] = r[i] - 13
			}
			str += string(r[i])
		}
		for _, k := range str {
			z01.PrintRune(k)
		}
		z01.PrintRune('\n')
	}
}
*/

/*
	rot 14
	Write a function rot14 that returns the string within the parameter transformed into a rot14 string. Each letter will be replaced by the letter 14 spots ahead in the alphabetical order.

	'z' becomes 'n' and 'Z' becomes 'N'. The case of the letter stays the same.
		$ go run .
		Vszzc! Vck ofs Mci?
		$
*/

/*
	package piscine

	import (
		"piscine"
		"github.com/01-edu/z01"
	)

	func main() {
		result := piscine.Rot14("Hello! How are You?")

		for _, r := range result {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}

	func Rot14(s string) string {
		ar := []rune(s)
		var str string
		for i := 0; i < len(ar); i++ {
			if ar[i] >= 'a' && ar[i] <= 'l' {
				ar[i] = ar[i] + 14
			} else if ar[i] >= 'A' && ar[i] <= 'L' {
				ar[i] = ar[i] + 14
			} else if ar[i] >= 'm' && ar[i] <= 'z' {
				ar[i] = ar[i] - 12
			} else if ar[i] >= 'M' && ar[i] <= 'Z' {
				ar[i] = ar[i] - 12
			}
			str += string(ar[i])
		}
		return str
}
*/
