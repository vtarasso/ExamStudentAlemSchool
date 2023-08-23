package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 || Atoi(args[0]) <= 1 {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}
	nb := Atoi(args[0])
	str := Addprimesum(nb)
	res := Itoa(str)
	for _, w := range res {
		z01.PrintRune(w)
	}
	z01.PrintRune('\n')
}

func Addprimesum(nb int) int {
	for nb <= 1 {
		return 0
	}
	for i := 2; i <= nb/2; i++ {
		if nb%i == 0 {
			return Addprimesum(nb - 1)
		}
	}
	return nb + Addprimesum(nb-1)
}

func Atoi(args string) int {
	s := 0
	temp := 1
	if args[0] == '-' {
		temp = -1
		args = args[1:]
	} else if args[0] == '+' {
		args = args[1:]
	}
	for _, w := range args {
		if w < '0' || w > '9' {
			os.Exit(0)
		}
		s = s * 10
		s = s + int(w-48)
	}
	return s * temp
}

func Itoa(n int) string {
	res := ""
	if n == 0 {
		res = ""
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

/*package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		arg := Atoi(os.Args[1])
		var a int
		for i := 1; i <= arg; i++ {
			if IsPrime(i) {
				a += i
			}
		}
		res := Itoa(a)
		for i := 0; i < len(res); i++ {
			z01.PrintRune(rune(res[i]))
		}
	} else {
		z01.PrintRune('0')
	}
	z01.PrintRune('\n')
}

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	if nb > 1 {
		for i := 2; i <= nb/2; i++ {
			if nb%i == 0 || nb == 1 {
				return false
			}
		}
	}
	return true
}

func Atoi(str string) int {
	ch := 1
	if str[0] == '-' {
		str = str[1:]
		ch = -1
	}
	num := 0
	for i := range str {
		num = num*10 + int(rune(str[i]-48))
	}
	return ch * num
}

func Itoa(n int) string {
	ch := ""
	if n < 0 {
		n = -n
		ch = "-"
	}
	digits := "0123456789"
	if n < 10 {
		return ch + digits[n:n+1]
	}
	return ch + Itoa(n/10) + digits[n%10:n%10+1]
}
*/
