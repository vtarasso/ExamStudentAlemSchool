package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		return
	} else {
		n, err := strconv.Atoi(args[0])
		if err == nil {
			for i := 1; i < 10; i++ {
				res := Itoa(i) + " x " + Itoa(n) + " = " + Itoa(i*n)
				for _, w := range res {
					z01.PrintRune(w)
				}
				z01.PrintRune('\n')
			}
		}
	}
}

func Itoa(n int) string {
	res := ""
	if n == 0 {
		res = "0"
	} else if n < 0 {
		for n != 0 {
			res = string(rune(-n%10 + 48)) + res
			n = n / 10
		}
		res = string('-') + res
	} else {
		for n != 0 {
			res = string(rune(n%10+48)) + res
			n = n / 10
		}
	}
	return res
}

/*
func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		return
	}
	s, err := strconv.Atoi(args[0])
	if err == nil {
		for i := 1; i < 10; i++ {
			result := strconv.Itoa(i) + " x " + strconv.Itoa(s) + "= " + strconv.Itoa(i*s)
			for _, w := range result {
				z01.PrintRune(w)
			}
			z01.PrintRune('\n')
		}
	}
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
	arg := os.Args[1:]
	if len(arg) != 1 {
		z01.PrintRune(10)
	} else {
		intn, err := strconv.Atoi(arg[0])
		if err == nil {
			if intn < 0 {
				z01.PrintRune(10)
				return
			}
			for i := 1; i < 10; i++ {
				valid(i)
				z01.PrintRune(' ')
				z01.PrintRune('x')
				z01.PrintRune(' ')
				valid(intn)
				z01.PrintRune(' ')
				z01.PrintRune('=')
				z01.PrintRune(' ')
				valid(i * intn)
				z01.PrintRune(10)
			}
		}
	}
}
func valid(n int) {
	if n >= 10 {
		valid(n / 10)
	}
	z01.PrintRune(rune(n%10 + 48))
}
*/

/*func main() {
	if len(os.Args) != 2 {
		z01.PrintRune('\n')
		return
	}
	arg, _ := strconv.Atoi(os.Args[1])
	for i := 1; i <= 9; i++ {
		Print(Itoa(i) + " " + "x" + " " + Itoa(arg) + " " + "=" + " " + Itoa(arg*i))
	}
}
func Print(str string) {
	for _, w := range str {
		z01.PrintRune(w)
	}
	z01.PrintRune('\n')
}
func Itoa(n int) string {
	var flag bool
	a := ""
	if n < 0 {
		n *= -1
		flag = true
	}
	for n != 0 {
		a = string(rune(n%10)+48) + a
		n = n / 10
	}
	if flag {
		a = "-" + a
	}
	return a
}*/

/*
package main

import (
	"os"
	"strconv"


	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		res, err := strconv.Atoi(os.Args[1])
		if err != nil {
			return
		}
		tabmult(res)
	} else {
		z01.PrintRune(10)
	}
}

func PrintString(s string) {
	for i := 0; i < len(s); i++ {
		z01.PrintRune(rune(s[i]))
	}
	z01.PrintRune('\n')
}

func tabmult(num int) {
	for i := 1; i < 10; i++ {
		PrintString(strconv.Itoa(i) + " x " + strconv.Itoa(num) + " = " + strconv.Itoa(i*num)) // 1 x 19 = 19
	}
}
*/
