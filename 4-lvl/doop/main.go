//точно рабочий)
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 4 {
		arg := os.Args[1:]
		a := Atoi(arg[0])
		b := Atoi(arg[2])
		
		if len(arg[0]) >= 18 || len(arg[2]) >= 18 {
			return
		}
		switch arg[1] {
		case "+":
			fmt.Printf("%d\n", a+b)
		case "-":
			fmt.Printf("%d\n", a-b)
		case "*":
			fmt.Printf("%d\n", a*b)
		case "/":
			if b == 0 {
				fmt.Printf("No division by 0\n")
				return
			}
			fmt.Printf("%d\n", a/b)
		case "%":
			if b == 0 {
				fmt.Printf("No modulo by 0\n")
				return
			}
			fmt.Printf("%d\n", a%b)
		}
	}
}

func Atoi(args string) int {
	s := 0 
	temp := 1
	if args[0] == '-' {
		temp = -1 
		args = args[1:]
	}
	
	for _, w := range args {
		if w < '0' || w > '9' { 
			os.Exit(0)
		}
		s = s*10 + int(w-48) 
	}
	return s * temp
}

/* perfect doop 
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 4 {
		args := os.Args[1:]
		a, aErr := Atoi(args[0])
		b, bErr := Atoi(args[2])

		if !aErr && !bErr {
			switch args[1] {
			case "+":
				c := a + b
				if (c > a) == (b > 0) {
					fmt.Printf("%d\n", c)
				}
			case "-":
				c := a - b
				if (c < a) == (b > 0) {
					fmt.Printf("%d\n", c)
				}
			case "*":
				c := a * b
				if c/a == b {
					fmt.Printf("%d\n", c)
				}
			case "/":
				if b == 0 {
					fmt.Printf("No division by 0\n")
					return
				}
				fmt.Printf("%d\n", a/b)
			case "%":
				if b == 0 {
					fmt.Printf("No modulo by 0\n")
					return
				}
				fmt.Printf("%d\n", a%b)
			default:
				return
			}
		}

	}
}

func Atoi(args string) (int, bool) {
	s := 0
	temp := 1
	if args[0] == '-' {
		temp = -1
		args = args[1:]
	}
	for _, w := range args {
		if w < '0' || w > '9' {
			os.Exit(0) // return
		}
		x := s * 10
		y := int(w - 48)
		t := x + y
		if (t > x) == (y > 0) {
			s = t
		} else {
			return -1, true
		}
	}
	return s * temp, false
}
*/

/*package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 4 {
		args := os.Args[1:]
		a := Atoi(args[0])
		b := Atoi(args[2])

		if a > 0 && b > 0 && (a+b) < 0 || a < 0 && b < 0 && (a+b) > 0 {
			return // os.Exit(0)
		}

		switch args[1] {
		case "+":
			c := a + b
			if (c > a) == (b > 0) {
				fmt.Printf("%d\n", c)
			}
		case "-":
			c := a - b
			if (c < a) == (b > 0) {
				fmt.Printf("%d\n", c)
			}
		case "*":
			c := a * b
			if c/a == b {
				fmt.Printf("%d\n", c)
			}
		case "/":
			if b == 0 {
				fmt.Printf("No division by 0\n")
				return
			}
			fmt.Printf("%d\n", a/b)
		case "%":
			if b == 0 {
				fmt.Printf("No modulo by 0\n")
				return
			}
			fmt.Printf("%d\n", a%b)
		default:
			return
		}
	}
}

func Atoi(args string) int {
	s := 0
	temp := 1
	if args[0] == '-' {
		temp = -1
		args = args[1:]
	}
	for _, w := range args {
		if w < '0' || w > '9' {
			os.Exit(0) // return
		}
		s = s*10 + int(w-48)
	}
	return s * temp
}
*/
/*package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 4 {
		arg := os.Args[1:]
		a := Atoi(arg[0])
		b := Atoi(arg[2])
		if a > 0 && b > 0 && (a+b) < 0 || a < 0 && b < 0 && (a+b) > 0 {
			return
		}
		if len(arg[0]) > 18 || len(arg[2]) > 18 { // if len(arg[0]) >= 19 || len(arg[2]) >= 19
			return
		}
		switch arg[1] {
		case "+":
			fmt.Printf("%d\n", a+b)
		case "-":
			fmt.Printf("%d\n", a-b)
		case "*":
			fmt.Printf("%d\n", a*b)
		case "/":
			if b == 0 {
				fmt.Printf("No division by 0\n")
				return
			}
			fmt.Printf("%d\n", a/b)
		case "%":
			if b == 0 {
				fmt.Printf("No modulo by 0\n")
				return
			}
			fmt.Printf("%d\n", a%b)
		default:
			return
		}
	}
}

func Atoi(str string) int {
	neg := 1 // var neg bool // for false
	if str[0] == '-' {
		neg = -1 // neg = true // for true
		str = str[1:]
	}
	n := 0 // var n int
	for _, w := range str {
		if w < '0' || w > '9' { // if w < 48 || w > 58
			os.Exit(0)
		}
		n = n*10 + int(w-'0') // n*10 + int(w) - 48
	}
	return neg * n
}
*/
/*package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 4 {
		return
	}
	a := Atoi(os.Args[1])
	b := Atoi(os.Args[3])
	op := os.Args[2]
	maxInt := int(^uint(0) >> 1)
	minInt := -maxInt - 1
	if (a > 0 && b > 0 && a+b < 0) || a > maxInt || b > maxInt || a < minInt || b < minInt {
		os.Exit(0)
	}
	switch op {
	case "+":
		c := a + b
		fmt.Printf("%d\n", c)
	case "-":
		c := a - b
		fmt.Printf("%d\n", c)
	case "*":
		c := a * b
		fmt.Printf("%d\n", c)
	case "/":
		if b == 0 {
			fmt.Printf("No division by 0\n")
			return
		}
		c := a / b
		fmt.Printf("%d\n", c)
	case "%":
		if b == 0 {
			fmt.Printf("No modulo by 0\n")
			return
		}
		c := a % b
		fmt.Printf("%d\n", c)
	}
}

func Atoi(args string) int {
	s := 0
	temp := 1
	for _, w := range args {
		if w == '-' {
			temp = -1
			args = args[1:]
		}
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
*/
/*package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		return
	}
	a := atoi(args[0])
	b := atoi(args[2])
	if a > 0 && b > 0 && (a+b) < 0 || (a < 0 && b < 0 && (a+b) > 0) {
		return
	}
	switch args[1] {
	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	case "*":
		fmt.Println(a * b)
	case "/":
		if b == 0 {
			fmt.Printf("No division by 0\n")
			return
		}
		fmt.Println(a / b)
	case "%":
		if b == 0 {
			fmt.Printf("No module by 0\n")
			return
		}
		fmt.Println(a % b)
	}
}

func atoi(str string) int {
	negative := 1
	if str[0] == '-' {
		negative = -1
		str = str[1:]
	}
	num := 0
	for _, w := range str {
		if w < '0' || w > '9' {
			os.Exit(0)
		}
		num = num*10 + int(w-'0')
	}
	return negative * num
}
*/

/*
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 4 {
		return
	}
	a := Atoi(os.Args[1])
	b := Atoi(os.Args[3])
	op := os.Args[2]
	maxInt := int(^uint(0) >> 1)
	minInt := -maxInt - 1
	if (a > 0 && b > 0 && a+b < 0) || a > maxInt || b > maxInt || a < minInt || b < minInt {
		os.Exit(0)
	}
	switch op {
	case "+":
		c := a + b
		fmt.Printf("%d\n", c)
	case "-":
		c := a - b
		fmt.Printf("%d\n", c)
	case "*":
		c := a * b
		fmt.Printf("%d\n", c)
	case "/":
		if b == 0 {
			fmt.Printf("No division by 0\n")
			return
		}
		c := a / b
		fmt.Printf("%d\n", c)
	case "%":
		if b == 0 {
			fmt.Printf("No modulo by 0\n")
			return
		}
		c := a % b
		fmt.Printf("%d\n", c)
	}
}

func Atoi(s string) int {
	var num int
	temp := 1
	if len(s) > 0 {
		if s[0] == '-' {
			s = s[1:]
			temp = -1
		}
		for _, w := range s {
			if w < '0' || w > '9' {
				os.Exit(0)
			}
			num = num*10 + int(w) - '0'
		}
	}
	return num * temp
}
*/

/*
package main
import (
	"os"
	"github.com/01-edu/z01"
)
func main() {
	if len(os.Args) != 4 {
		return
	}
	operator := os.Args[2]
	v1 := os.Args[1]
	v2 := os.Args[3]
	n1 := atoi(v1)
	n2 := atoi(v2)
	var result int
	if !checkOpetator(operator) || !checkNumber(v1) || !checkNumber(v2) || !isValid(n1) || !isValid(n2) || len(v1) < 1 || len(v2) < 1 {
		return
	}
	switch operator {
	case "+":
		result = n1 + n2
	case "-":
		result = n1 - n2
	case "/":
		if n2 == 0 {
			printStr("No division by 0")
			return
		}
		result = n1 / n2
	case "*":
		result = n1 * n2
	case "%":
		if n2 == 0 {
			printStr("No modulo by 0")
			return
		}
		result = n1 % n2
	}
	if !isValid(result) {
		return
	}
	printN(result)
}
func isValid(n int) bool {
	if n >= -2147483648 && n <= 2147483647 {
		return true
	}
	return false
}
func checkNumber(s string) bool {
	for _, r := range s {
		if r >= '0' && r <= '9' {
			return true
		}
	}
	return false
}
func checkOpetator(s string) bool {
	if s == "+" || s == "-" || s == "/" || s == "*" || s == "%" {
		return true
	}
	return false
}
func atoi(s string) int {
	s0 := s
	if s[0] == '-' || s[0] == '+' {
		s = s[1:]
	}
	n := 0
	for _, ch := range []byte(s) {
		ch -= '0'
		if ch > 9 {
			return 0
		}
		n = n*10 + int(ch)
	}
	if s0[0] == '-' {
		n = -n
	}
	return n
}
func printN(a int) {
	if a == 0 {
		z01.PrintRune('0')
	} else if a < 0 {
		z01.PrintRune('-')
	}
	slice := []int{}
	for a != 0 {
		slice = append(slice, a%10)
		a /= 10
	}
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] < 0 {
			slice[i] *= -1
		}
		z01.PrintRune(rune(slice[i] + 48))
	}
	z01.PrintRune('\n')
}
func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
*/
/*package main
import (
	"os"
	"github.com/01-edu/z01"
)
func PrintStr(s string) {
	sByByte := []byte(s)
	for _, v := range sByByte {
		z01.PrintRune(rune(v))
	}
	z01.PrintRune('\n')
}
func PrintNbr(n int) {
	var arrayOfInts []int
	if n == 0 {
		arrayOfInts = append(arrayOfInts, 0)
	} else if n < 0 {
		for n != 0 {
			arrayOfInts = append(arrayOfInts, -(n % 10))
			n /= 10
		}
		z01.PrintRune('-')
	} else {
		for n != 0 {
			arrayOfInts = append(arrayOfInts, n%10)
			n /= 10
		}
	}
	for i := len(arrayOfInts) - 1; i >= 0; i-- {
		z01.PrintRune(rune(arrayOfInts[i] + 48))
	}
	z01.PrintRune('\n')
}
func IsValidValue(s string) bool {
	for _, v := range s {
		if (v < '0' || v > '9') && v != '-' {
			return false
		}
	}
	return true
}
func IsValidOperator(s string) bool {
	for _, v := range s {
		if v == '+' || v == '-' || v == '/' || v == '*' || v == '%' {
			return true
		}
	}
	return false
}
func Calculate(f func(s string) (int, bool), first, operator, second string) {
	a, aBool := f(first)
	b, bBool := f(second)
	if aBool && bBool {
		switch operator {
		case "+":
			c := a + b
			if (c > a) == (b > 0) {
				PrintNbr(c)
			}
		case "-":
			c := a - b
			if (c < a) == (b > 0) {
				PrintNbr(c)
			}
		case "*":
			c := a * b
			if (c < 0) == ((a < 0) != (b < 0)) {
				if c/b == a {
					PrintNbr(c)
				}
			}
		case "/":
			if b != 0 {
				c := a / b
				PrintNbr(c)
			} else {
				PrintStr("No division by 0")
			}
		case "%":
			if b != 0 {
				c := a % b
				PrintNbr(c)
			} else {
				PrintStr("No modulo by 0")
			}
		}
	}
}
func Atoi(s string) (int, bool) {
	sum := 0
	if len(s) > 0 {
		for i, v := range s {
			if i == 0 && v == '+' {
			} else if i == 0 && v == '-' {
			} else if v >= '0' && v <= '9' {
				c := sum*10 + int(v-48)
				if (c > sum*10) == (int(v-48) > 0) {
					sum = c
				} else {
					return 0, false
				}
			} else {
				return 0, false
			}
		}
		if s[0] == '-' {
			return -sum, true
		} else {
			return sum, true
		}
	}
	return 0, false
}
func main() {
	args := os.Args[1:]
	if len(args) == 3 && IsValidValue(args[0]) && IsValidOperator(args[1]) && IsValidValue(args[2]) {
		Calculate(Atoi, args[0], args[1], args[2])
	}
}
*/

/*
package main
import (
	"os"
	"github.com/01-edu/z01"
)
func PrintStr(s string) {
	sByByte := []byte(s)
	for _, v := range sByByte {
		z01.PrintRune(rune(v))
	}
	z01.PrintRune('\n')
}
func PrintNbr(n int) {
	var arrayOfInts []int
	if n == 0 {
		arrayOfInts = append(arrayOfInts, 0)
	} else if n < 0 {
		for n != 0 {
			arrayOfInts = append(arrayOfInts, -(n % 10))
			n /= 10
		}
		z01.PrintRune('-')
	} else {
		for n != 0 {
			arrayOfInts = append(arrayOfInts, n%10)
			n /= 10
		}
	}
	for i := len(arrayOfInts) - 1; i >= 0; i-- {
		z01.PrintRune(rune(arrayOfInts[i] + 48))
	}
	z01.PrintRune('\n')
}
func IsValidValue(s string) bool {
	for _, v := range s {
		if (v < '0' || v > '9') && v != '-' {
			return false
		}
	}
	return true
}
func IsValidOperator(s string) bool {
	for _, v := range s {
		if v == '+' || v == '-' || v == '/' || v == '*' || v == '%' {
			return true
		}
	}
	return false
}
func Calculate(f func(s string) (int, bool), first, operator, second string) {
	a, aBool := f(first)
	b, bBool := f(second)
	if aBool && bBool {
		switch operator {
		case "+":
			c := a + b
			if (c > a) == (b > 0) {
				PrintNbr(c)
			}
		case "-":
			c := a - b
			if (c < a) == (b > 0) {
				PrintNbr(c)
			}
		case "*":
			c := a * b
			if (c < 0) == ((a < 0) != (b < 0)) {
				if c/b == a {
					PrintNbr(c)
				}
			}
		case "/":
			if b != 0 {
				c := a / b
				PrintNbr(c)
			} else {
				PrintStr("No division by 0")
			}
		case "%":
			if b != 0 {
				c := a % b
				PrintNbr(c)
			} else {
				PrintStr("No modulo by 0")
			}
		}
	}
}
func Atoi(s string) (int, bool) {
	sum := 0
	if len(s) > 0 {
		for i, v := range s {
			if i == 0 && v == '+' {
			} else if i == 0 && v == '-' {
			} else if v >= '0' && v <= '9' {
				c := sum*10 + int(v-48)
				if (c > sum*10) == (int(v-48) > 0) {
					sum = c
				} else {
					return 0, false
				}
			} else {
				return 0, false
			}
		}
		if s[0] == '-' {
			return -sum, true
		} else {
			return sum, true
		}
	}
	return 0, false
}
func main() {
	args := os.Args[1:]
	if len(args) == 3 && IsValidValue(args[0]) && IsValidOperator(args[1]) && IsValidValue(args[2]) {
		Calculate(Atoi, args[0], args[1], args[2])
	}
}

*/
/*package main
import (
	"fmt"
	"os"
)
func main() {
	if len(os.Args) != 4 {
		return
	}
	arg := os.Args[1:]
	a := Atoi(arg[0])
	b := Atoi(arg[2])
	if a > 0 && b > 0 && (a+b) < 0 || a < 0 && b < 0 && (a+b) > 0 {
		return
	}
	if len(arg[0]) >= 19 {
		return
	}
	switch arg[1] {
	case "+":
		fmt.Printf("%d\n", a+b)
	case "-":
		fmt.Printf("%d\n", a-b)
	case "*":
		fmt.Printf("%d\n", a*b)
	case "/":
		if b == 0 {
			fmt.Printf("No division by 0\n")
			return
		}
		fmt.Printf("%d\n", a/b)
	case "%":
		if b == 0 {
			fmt.Printf("No module by 0\n")
			return
		}
		fmt.Printf("%d\n", a%b)
	}
}
func Atoi(str string) int {
	neg := 1 // var neg bool // for false
	if str[0] == '-' {
		neg = -1 // neg = true // for true
		str = str[1:]
	}
	n := 0 // var n int
	for _, w := range str {
		if w < '0' || w > '9' { // if w < 48 || w > 58
			os.Exit(0)
		}
		n = n*10 + int(w-'0') // n*10 + int(w) - 48
	}
	return neg * n
}*/
