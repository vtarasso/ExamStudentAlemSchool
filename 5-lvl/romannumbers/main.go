package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		return
	}
	num := Atoi(args[0])
	if num <= 0 || num >= 4000 {
		fmt.Printf("ERROR: cannot convert to roman digit")
		fmt.Printf("\n")
		return
	}
	ones := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	tens := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	arns := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	thns := []string{"", "M", "MM", "MMM"}
	res := thns[num/1000] + arns[num%1000/100] + tens[num%100/10] + ones[num%10]
	ones1 := []string{"", "I+", "I+I+", "I+I+I+", "(V-I)+", "V+", "V+I+", "V+I+I+", "V+I+I+I+", "(X-I)+"}
	tens1 := []string{"", "X+", "X+X+", "X+X+X+", "(L-X)+", "L+", "L+X+", "L+X+X+", "L+X+X+X+", "(C-X)+"}
	arns1 := []string{"", "C+", "C+C+", "C+C+C+", "(D-C)+", "D+", "D+C+", "D+C+C+", "D+C+C+C+", "(M-C)+"}
	thns1 := []string{"", "M+", "M+M+", "M+M+M+"}
	res1 := thns1[num/1000] + arns1[num%1000/100] + tens1[num%100/10] + ones1[num%10]
	fmt.Printf(res1[:len(res1)-1])
	fmt.Printf("\n")
	fmt.Printf(res)
	fmt.Printf("\n")
}

func Atoi(args string) int {
	s := 0
	temp := 1
	if args[0] == '-' {
		temp = -1
		args = args[1:]
	}
	for _, v := range args {
		if v < '0' || v > '9' {
			return 0
		}
		s = s * 10
		s = s + int(v-48)
	}
	return s * temp
}

/*package main
import (
	"fmt"
	"os"
)
func main() {
	if len(os.Args) == 2 {
		args := os.Args[1:]
		n := 0
		for _, w := range args[0] {
			if w < '0' || w > '9' {
				fmt.Printf("ERROR: cannot convert to roman digit\n")
				return
			}
			n = n*10 + int(w-'0')
		}
		if n >= 4000 {
			fmt.Printf("ERROR: cannot convert to roman digit\n")
			return
		}
		num := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
		romnum := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
		mathnum := []string{"M", "(M-C)", "D", "(D-C)", "C", "(C-X)", "L", "(L-X)", "X", "(X-I)", "V", "(V-I)", "I"}
		str := ""
		sum := ""
		for i := range num {
			temp := n / num[i]
			for temp > 0 {
				n -= num[i]
				str += romnum[i]
				sum += mathnum[i] + "+"
				temp--
			}
		}
		fmt.Printf("%s\n%s\n", sum[:len(sum)-1], str)
	}
}*/
